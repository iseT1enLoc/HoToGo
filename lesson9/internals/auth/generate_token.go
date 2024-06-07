package auth

import (
	usermodel "lesson9/internals/modules/user/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user usermodel.User) (string, error) {
	claims := &jwt.MapClaims{
		"sub":        user.ID,
		"exp":        time.Now().Add(time.Hour * 24 * 30).Unix(),
		"authorized": true,
	}
	//generate the jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//sign and get complete encoded token
	tokenstring, err := token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenstring, err
}
