package userhandlers

import (
	"fmt"
	"lesson9/internals/auth"
	appcontext1 "lesson9/internals/common/appcontext"
	usermodel "lesson9/internals/modules/user/models"
	"lesson9/internals/templates"
	"lesson9/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(appctx appcontext1.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//bind the request body
		if c.Request.Method == "POST" {
			//1.get email from request body
			var body struct {
				Email    string
				Password string
			}
			/*if c.BindJSON(&body) != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Can not bind user sign-in body"})
				return
			}*/
			body.Email = c.Request.FormValue("email")
			body.Password = c.Request.FormValue("password")

			//2.look up requested user
			var user usermodel.User
			db := appctx.GetConnectionToDatabase()
			db.First(&user, "email = ?", body.Email)
			fmt.Printf("Email: %v,Password: %v\n", body.Email, body.Password)
			fmt.Printf("Email: %v,Password: %v\n", user.Email, user.Password)
			if user.ID == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or Password"})
				return
			}
			//3.compare sent in pass with saved users pass hash
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong password"})
				return
			}
			//4.generate the jwt token
			tokenstring, err := auth.GenerateToken(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to create token"})
				return
			}
			//5.send it back
			//c.SetSameSite(http.SameSiteLaxMode)
			//c.SetCookie("Authorization", tokenstring, 3600*24*30, "", "", false, true)
			expirationTime := time.Now().Add(5 * time.Minute)
			c.SetCookie("token", tokenstring, int(expirationTime.Unix()), "/", "localhost", false, true)
			c.Redirect(http.StatusFound, "/index")
			return
		}

		//c.JSON(http.StatusOK, gin.H{"token": tokenstring})
		utils.Render(c, http.StatusOK, templates.LoginPage())
	}
}
