package userhandlers

import (
	"fmt"
	"lesson8/component/appcontext"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignOut(appctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Enter signout")
		// Delete the session
		session := sessions.Default(c)
		fmt.Println("line17")
		session.Clear()
		fmt.Println("line19")
		session.Save()
		fmt.Println("line21")
		c.Status(http.StatusAccepted)
	}
}
