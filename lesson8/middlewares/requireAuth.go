package middlewares

import (
	"fmt"
	"lesson8/component/appcontext"
	usermodel "lesson8/modules/users/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequiredAuth(appctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get the cookie of request
		tokenString, err := c.Cookie("Authorization")

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//decode validate it
		// Parse the token and verify its signature
		fmt.Println("Entered line 25")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check that the signing method is what we expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the secret key
			return []byte(os.Getenv("SECRET")), nil
		})
		fmt.Println("Entered line 35")
		// Check for errors
		if err != nil {
			//http.StatusUnauthorized == 401
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		fmt.Println("Entered line 41")
		// Check that the token is valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("Entered line 44")
			//check the expire
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			//Find the user with token sub
			var user usermodel.User
			db := appctx.GetConnectionToDatabase()
			db.First(&user, claims["sub"])
			if user.ID == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or Password"})
				return
			}
			//Attach the request
			c.Set("user", user)

			//Continue
			c.Next()
		} else {
			fmt.Println("Entered line 64")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
