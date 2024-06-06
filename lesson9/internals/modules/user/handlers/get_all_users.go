package userhandlers

import (
	appcontext1 "lesson9/internals/common/appcontext"
	userstorage "lesson9/internals/modules/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(appctx appcontext1.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userstorage.ListUser(appctx.GetConnectionToDatabase())
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, users)
	}
}
