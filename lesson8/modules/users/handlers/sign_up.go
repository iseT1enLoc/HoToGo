package userhandlers

import (
	"lesson8/component/appcontext"
	usermodel "lesson8/modules/users/model"
	userstorage "lesson8/modules/users/storage"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(appctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get email from request body
		var reqCreateUser usermodel.ReqCreateUser
		if err := c.ShouldBindJSON(&reqCreateUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Can not bind user sign-up body"})
			return
		}

		//hash the password
		hash, err := bcrypt.GenerateFromPassword([]byte(reqCreateUser.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Can not hash password"})
			return
		}
		//create the user
		user := usermodel.User{
			Name:     reqCreateUser.Name,
			Email:    reqCreateUser.Email,
			Password: string(hash),
		}
		err = userstorage.CreateUser(appctx.GetConnectionToDatabase(), &user)
		if err != nil {
			panic(err)
		}
		// Set the session
		session := sessions.Default(c)
		session.Set("userID", user.ID)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save session"})
			return
		}
		c.JSON(http.StatusCreated, user)
	}
}
