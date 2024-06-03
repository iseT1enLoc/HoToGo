package userhandlers

import (
	"lesson8/component/appcontext"
	userstorage "lesson8/modules/users/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUser(appctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userstorage.ListUser(appctx.GetConnectionToDatabase())
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, users)
	}
}
