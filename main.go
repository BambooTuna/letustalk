package main

import (
	"database/sql"
	"fmt"
	"github.com/BambooTuna/letustalk/backend/application"
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/infrastructure/persistence"
	"github.com/BambooTuna/letustalk/backend/interfaces"
	"github.com/BambooTuna/letustalk/backend/json"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"log"
	"net/http"
	"os"
)
func main() {

	mysqlDataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.FetchEnvValue("MYSQL_USER", "BambooTuna"),
		config.FetchEnvValue("MYSQL_PASS", "pass"),
		config.FetchEnvValue("MYSQL_HOST", "127.0.0.1"),
		config.FetchEnvValue("MYSQL_PORT", "3306"),
		config.FetchEnvValue("MYSQL_DATABASE", "letustalk"),
	)
	db, err := sql.Open("mysql", mysqlDataSourceName)
	dbSession := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbSession.Db.Close()
	if err != nil {
		log.Fatal(err)
	}

	invoiceDetailRepository := persistence.InvoiceDetailRepositoryImpl{DBSession:dbSession}
	invoiceDetailUseCase := application.InvoiceDetailUseCase{InvoiceDetailRepository:invoiceDetailRepository}
	invoiceDetailHandler := interfaces.InvoiceDetailHandler{InvoiceDetailUseCase:invoiceDetailUseCase}

	apiVersion := "/v1"

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./front/dist", false)))

	api := r.Group(apiVersion)

	api.GET("/invoice/:invoiceId", invoiceDetailHandler.UnimplementedRoute())
	api.POST("/invoice", invoiceDetailHandler.UnimplementedRoute())

	api.POST("/pay/:invoiceId", invoiceDetailHandler.UnimplementedRoute())

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
