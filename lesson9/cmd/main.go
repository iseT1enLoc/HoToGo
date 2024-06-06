package main

import (
	appcontext1 "lesson9/internals/common/appcontext"
	"lesson9/internals/db"
	middlewares "lesson9/internals/middleware"
	userhandlers "lesson9/internals/modules/user/handlers"
	"lesson9/internals/templates"
	"lesson9/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := db.LoadConfig()
	if err != nil {
		log.Fatalf("Can not load config", err)
		return
	}
	db, err := db.GetConnectionToDatabaseInboundedTime(cfg)
	appctx := appcontext1.NewAppCtx(db)

	router := gin.Default()
	router.Use(middlewares.CORS())
	router.Use(middlewares.Recover(appctx))

	//sign in handlers
	router.GET("/signin", userhandlers.SignIn(appctx))
	router.POST("/signin", userhandlers.SignIn(appctx))

	//sign up handlers
	router.GET("/signup", userhandlers.SignUpUser(appctx))
	router.POST("/signup", userhandlers.SignUpUser(appctx))

	router.DELETE("/signout", userhandlers.SignOut(appctx))

	router.GET("/index", middlewares.RequiredAuth(appctx), func(c *gin.Context) {
		utils.Render(c, http.StatusOK, templates.Index())
	})
	router.GET("/hello", middlewares.RequiredAuth(appctx), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "everything ok"})
	})

	/*public := router.Group("/public")
	{
		public.POST("/signup", userhandlers.SignUpUser(appctx))
		public.POST("/signin", userhandlers.SignIn(appctx))
	}
	requireAuth := router.Group("/requireAuth")
	requireAuth.Use(middlewares.RequiredAuth(appctx))
	{
		requireAuth.GET("/validate")
	}*/

	router.Run(":8080")
}
