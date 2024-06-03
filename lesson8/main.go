package main

import (
	"fmt"
	"lesson8/component/appcontext"
	"lesson8/config"
	"lesson8/middlewares"
	userhandlers "lesson8/modules/users/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Can not load config")
		return
	}
	fmt.Println("Connecting to database......")
	db, err := config.ConnectDatabaseInBoundedTime(cfg)
	if err != nil {
		log.Fatalln("Error when connecting to database:", err)
	}
	appctx := appcontext.NewAppCtx(db)

	router := gin.Default()
	router.Use(middlewares.CORS())
	router.Use(middlewares.Recover(appctx))

	router.POST("/users", userhandlers.CreateUser(appctx))
	router.GET("/users", userhandlers.ListUser(appctx))
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalln("Error running server:", err)
	}
}
