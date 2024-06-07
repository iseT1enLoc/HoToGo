package userhandlers

import (
	appcontext1 "lesson9/internals/common/appcontext"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(appctx appcontext1.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")

		c.JSON(http.StatusOK, gin.H{"message": user})
	}
}
