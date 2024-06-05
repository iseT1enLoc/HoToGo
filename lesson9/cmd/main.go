package main

import (
	"lesson9/internals/templates"
	"lesson9/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		utils.Render(c, http.StatusOK, templates.Hello("Hello"))
	})
	r.Run(":8080")
}
