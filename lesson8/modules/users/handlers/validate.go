package userhandlers

import (
	"lesson8/component/appcontext"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(appctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		c.JSON(http.StatusOK, gin.H{"message": user})
	}
}
