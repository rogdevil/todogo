package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rogdevil/todo/config"
	"github.com/rogdevil/todo/migration"
	"github.com/rogdevil/todo/route"
)

func init() {
	db := config.Init()
	migration.Migrate(db)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := route.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
