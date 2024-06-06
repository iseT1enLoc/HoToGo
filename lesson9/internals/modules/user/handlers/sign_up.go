package userhandlers

import (
	"fmt"
	appcontext1 "lesson9/internals/common/appcontext"
	usermodel "lesson9/internals/modules/user/models"
	userstorage "lesson9/internals/modules/user/storage"
	"lesson9/internals/templates"
	"lesson9/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(appctx appcontext1.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			//1.Bind user reqCreate
			//var reqCreateUser usermodel.ReqUser
			//1.1 for api calling
			/*if err := c.ShouldBindJSON(&reqCreateUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Can not bind user sign-up body"})
				return
			}*/
			//1.2 for ui communication
			reqCreateUser := usermodel.ReqUser{
				UserName: c.Request.FormValue("username"),
				Email:    c.Request.FormValue("email"),
				Password: c.Request.FormValue("password"),
			}
			fmt.Printf("Get data successfully")
			fmt.Printf("%v %v %v\n\n", reqCreateUser.UserName, reqCreateUser.Email, reqCreateUser.Password)
			//2.hash the password
			hash, err := bcrypt.GenerateFromPassword([]byte(reqCreateUser.Password), 10)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Can not hash password"})
				return
			}
			//3.create the user
			user := usermodel.User{
				UserName: reqCreateUser.UserName,
				Email:    reqCreateUser.Email,
				Password: string(hash),
			}
			//4.InsertDB
			err = userstorage.AddUser(appctx.GetConnectionToDatabase(), &user)
			if err != nil {
				panic(err)
			}
			fmt.Print("Get to line 51\n")
			c.Redirect(http.StatusFound, "/hello")
			//redirect after user created
			//utils.Render(c, http.StatusCreated, templates.Hello("Hello"))
			return
		}
		fmt.Printf("Enter line 55\n\n")
		//get request
		utils.Render(c, http.StatusOK, templates.SignUp())
	}
}
