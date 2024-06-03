package userhandlers

import (
	"fmt"
	"lesson8/component/appcontext"
	usermodel "lesson8/modules/users/model"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(appctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//bind the request body
		//get email from request body
		var body struct {
			Email    string
			Password string
		}
		if c.BindJSON(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Can not bind user sign-in body"})
			return
		}
		//look up requested user
		var user usermodel.User
		db := appctx.GetConnectionToDatabase()
		db.First(&user, "email = ?", body.Email)
		fmt.Printf("Email: %v,Password: %v\n", body.Email, body.Password)
		fmt.Printf("Email: %v,Password: %v\n", user.Email, user.Password)
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or Password"})
			return
		}
		//compare sent in pass with saved users pass hash
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong password"})
			return
		}
		//generate the jwt token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": user.ID, "exp": time.Now().Add(time.Hour * 24 * 30).Unix()})
		//sign and get complete encoded token
		tokenstring, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to create token"})
			return
		}
		//send it back
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenstring, 3600*24*30, "", "", false, true)
		c.JSON(http.StatusOK, gin.H{"token": tokenstring})
	}
}
