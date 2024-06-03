package userhandlers

import (
	"lesson8/component/appcontext"
	usermodel "lesson8/modules/users/model"
	userstorage "lesson8/modules/users/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(appctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqCreateUser usermodel.ReqCreateUser
		if err := c.ShouldBindJSON(&reqCreateUser); err != nil {
			panic(err)
		}

		user := usermodel.User{
			Name: reqCreateUser.Name,
		}
		err := userstorage.CreateUser(appctx.GetConnectionToDatabase(), &user)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusCreated, user)
	}
}
