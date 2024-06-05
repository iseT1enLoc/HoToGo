package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Successfully log in")
	})
	r.Run(":8080")
}
