package main

import (
	"fmt"
	"github.com/BambooTuna/go-server-lib/session"
	"github.com/BambooTuna/letustalk/backend/application"
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/BambooTuna/letustalk/backend/infrastructure"
	"github.com/BambooTuna/letustalk/backend/infrastructure/persistence"
	"github.com/BambooTuna/letustalk/backend/interfaces"
	"github.com/BambooTuna/letustalk/backend/interfaces/json"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/payjp/payjp-go/v1"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	sessionRedisSession := infrastructure.RedisConnect(0)
	sessionDao := session.RedisSessionStorageDao{Client: sessionRedisSession}
	authSession := session.DefaultSession{Dao: sessionDao, Settings: session.DefaultSessionSettings(config.FetchEnvValue("SESSION_SECRET", "1234567890asdfghjkl"))}

	dbSession := infrastructure.GormConnect()
	defer dbSession.Close()

	pay := payjp.New("sk_test_140a9e4c676a5befdf04206e", nil)
	accountDetailRepository := persistence.AccountDetailRepositoryImpl{DBSession: dbSession}
	invoiceDetailRepository := persistence.InvoiceRepositoryImpl{DBSession: dbSession}
	reservationRepository := persistence.ReservationRepositoryImpl{DBSession: dbSession}
	scheduleRepository := persistence.ScheduleRepositoryImpl{DBSession: dbSession}

	accountDetailUseCase := application.AccountDetailUseCase{AccountDetailRepository: accountDetailRepository}
	invoiceDetailUseCase := application.InvoiceUseCase{InvoiceRepository: invoiceDetailRepository, PaymentService: pay}
	scheduleUseCase := application.ScheduleUseCase{ScheduleRepository: scheduleRepository, ReservationRepository: reservationRepository, InvoiceRepository: invoiceDetailRepository}

	accountDetailHandler := interfaces.AccountDetailHandler{Session: authSession, AccountDetailUseCase: accountDetailUseCase}
	invoiceDetailHandler := interfaces.InvoiceHandler{Session: authSession, InvoiceUseCase: invoiceDetailUseCase}
	scheduleHandler := interfaces.ScheduleHandler{Session: authSession, ScheduleUseCase: scheduleUseCase}

	apiVersion := "/v1"

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./front/dist", false)))

	api := r.Group(apiVersion)

	api.GET("/mentor", accountDetailHandler.GetMentorAccountDetailsRoute())
	api.GET("/account/:accountId", accountDetailHandler.GetAccountDetailRoute("accountId"))

	api.GET("/account/:accountId/schedule", scheduleHandler.GetFreeScheduleRoute("accountId"))

	api.POST("/schedule/:scheduleId/reserve", scheduleHandler.ReserveRoute("scheduleId"))

	api.GET("/invoice/:invoiceId", invoiceDetailHandler.GetInvoiceRoute("invoiceId"))
	api.POST("/invoice", invoiceDetailHandler.IssueAnInvoiceRoute())

	api.POST("/pay/:invoiceId", invoiceDetailHandler.MakePaymentRoute("invoiceId"))

	api.GET("/test", DBTestRoute(scheduleRepository.ResolveByParentAccountId("1", time.Now(), time.Now())))
	api.GET("/health", UnimplementedRoute)

	r.NoRoute(func(c *gin.Context) {
		c.File("./front/dist/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}

func UnimplementedRoute(ctx *gin.Context) {

	ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: "UnimplementedRoute"})
}

func DBTestRoute(result []*domain.Schedule) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, result)
	}
}
