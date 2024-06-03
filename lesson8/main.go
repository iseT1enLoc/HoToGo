package main

import (
	"fmt"
	"lesson8/component/appcontext"
	"lesson8/config"
	"lesson8/middlewares"
	userhandlers "lesson8/modules/users/handlers"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
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
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("app", store))
	router.Use(middlewares.CORS())
	router.Use(middlewares.Recover(appctx))

	router.POST("/users/signup", userhandlers.SignUp(appctx))
	router.POST("/users/signin", userhandlers.SignIn(appctx))
	router.GET("/validate", middlewares.RequiredAuth(appctx), userhandlers.Validate(appctx))
	router.DELETE("/users/signout", userhandlers.SignOut(appctx))
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalln("Error running server:", err)
	}
}
