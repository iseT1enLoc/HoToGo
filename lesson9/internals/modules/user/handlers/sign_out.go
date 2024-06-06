package userhandlers

import (
	"fmt"
	appcontext1 "lesson9/internals/common/appcontext"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignOut(appctx appcontext1.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Enter signout")
		// Delete the session
		/*session := sessions.Default(c)
		fmt.Println("line17")
		session.Clear()
		fmt.Println("line19")
		session.Save()
		fmt.Println("line21")*/
		c.SetCookie("Authorization", "", -1, "", "", false, true)
		//c.SetCookie("token", "", -1, "/", "localhost", false, true)
		c.Redirect(http.StatusOK, "/signin")
	}
}
