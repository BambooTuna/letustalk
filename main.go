package main

import (
	"fmt"
	"github.com/BambooTuna/letustalk/backend/json"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	apiVersion := "/v1"

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./front/dist", false)))

	r.GET(apiVersion+"/health", UnimplementedRoute)

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
