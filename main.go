package main

import (
	"fmt"
	"github.com/BambooTuna/go-server-lib/authentication"
	"github.com/BambooTuna/go-server-lib/session"
	"github.com/BambooTuna/letustalk/backend/application"
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/infrastructure"
	"github.com/BambooTuna/letustalk/backend/infrastructure/persistence"
	"github.com/BambooTuna/letustalk/backend/interfaces"
	"github.com/BambooTuna/letustalk/backend/interfaces/json"
	"github.com/BambooTuna/letustalk/docs"
	_ "github.com/BambooTuna/letustalk/docs"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/payjp/payjp-go/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
)

// @title Swagger Letustalk API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	port := config.FetchEnvValue("PORT", "8080")
	apiVersion := "/v1"

	sessionRedisSession := infrastructure.RedisConnect(0)
	activateRedisSession := infrastructure.RedisConnect(1)

	sessionDao := session.RedisSessionStorageDao{Client: sessionRedisSession}
	activateDao := session.RedisSessionStorageDao{Client: activateRedisSession}

	authSession := session.DefaultSession{Dao: sessionDao, Settings: session.DefaultSessionSettings(config.FetchEnvValue("SESSION_SECRET", "1234567890asdfghjkl"))}

	activateMailer := application.ActivateAccountMailerFromConfig()
	activatorUseCase := authentication.DefaultActivatorUseCase(activateDao, activateMailer)

	dbSession := infrastructure.GormConnect()
	defer dbSession.Close()

	pay := payjp.New("sk_test_140a9e4c676a5befdf04206e", nil)
	accountCredentialsRepository := persistence.AccountCredentialsRepositoryImpl{DBSession: dbSession}
	accountDetailRepository := persistence.AccountDetailRepositoryImpl{DBSession: dbSession}
	invoiceDetailRepository := persistence.InvoiceRepositoryImpl{DBSession: dbSession}
	reservationRepository := persistence.ReservationRepositoryImpl{DBSession: dbSession}
	scheduleRepository := persistence.ScheduleRepositoryImpl{DBSession: dbSession}

	accountCredentialsUseCase := application.AccountCredentialsUseCase{AccountCredentialsRepository: accountCredentialsRepository, ActivatorUseCase: activatorUseCase}
	accountDetailUseCase := application.AccountDetailUseCase{AccountDetailRepository: accountDetailRepository}
	invoiceDetailUseCase := application.InvoiceUseCase{InvoiceRepository: invoiceDetailRepository, PaymentService: pay}
	scheduleUseCase := application.ScheduleUseCase{AccountCredentialsRepository: accountCredentialsRepository, ScheduleRepository: scheduleRepository, ReservationRepository: reservationRepository, InvoiceRepository: invoiceDetailRepository}

	accountCredentialsHandler := interfaces.AccountCredentialsHandler{Session: authSession, AccountCredentialsUseCase: accountCredentialsUseCase}
	accountDetailHandler := interfaces.AccountDetailHandler{Session: authSession, AccountDetailUseCase: accountDetailUseCase}
	invoiceDetailHandler := interfaces.InvoiceHandler{Session: authSession, InvoiceUseCase: invoiceDetailUseCase}
	scheduleHandler := interfaces.ScheduleHandler{Session: authSession, ScheduleUseCase: scheduleUseCase}

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./front/dist", false)))

	api := r.Group(apiVersion)

	auth := api.Group("auth")
	auth.POST("/signup", accountCredentialsHandler.SignUpRoute())
	auth.POST("/signin", accountCredentialsHandler.SignInRoute())

	activate := api.Group("activate")
	activate.GET("/account/:code", accountCredentialsHandler.ActivateAccountRoute("code"))
	activate.PUT("/account", accountCredentialsHandler.SendActivateMailRoute())

	accounts := api.Group("accounts")
	//accounts.GET("/:accountId", accountDetailHandler.GetAccountDetailRoute("accountId"))
	accounts.GET("/:accountId/schedules", scheduleHandler.GetFreeScheduleRoute("accountId"))

	mentor := api.Group("mentor")
	mentor.GET("/", accountDetailHandler.GetMentorAccountDetailsRoute())

	reservations := api.Group("reservations")
	reservations.GET("/reserved/parent", scheduleHandler.GetReservedReservationsByParentAccountIdRoute())
	reservations.GET("/reserved/child", scheduleHandler.GetReservedReservationsByChildAccountIdRoute())
	reservations.POST("/reserve/:scheduleId", scheduleHandler.ReserveRoute("scheduleId"))

	invoices := api.Group("invoices")
	invoices.GET("/:invoiceId", invoiceDetailHandler.GetInvoiceRoute("invoiceId"))
	invoices.POST("/", invoiceDetailHandler.IssueAnInvoiceRoute())
	invoices.POST("/:invoiceId", invoiceDetailHandler.MakePaymentRoute("invoiceId"))

	api.GET("/health", UnimplementedRoute)

	docs.SwaggerInfo.Schemes = []string{config.FetchEnvValue("SWAGGER_SCHEMES", "http")}
	docs.SwaggerInfo.Host = config.FetchEnvValue("SWAGGER_HOST", fmt.Sprintf("localhost:%s", port))
	docs.SwaggerInfo.BasePath = apiVersion
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.NoRoute(func(c *gin.Context) {
		c.File("./front/dist/index.html")
	})

	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}

func UnimplementedRoute(ctx *gin.Context) {

	ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: "UnimplementedRoute"})
}
