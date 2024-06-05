package userhandlers

import (
	"lesson8/component/appcontext"
	userstorage "lesson8/modules/users/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUser(appctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userstorage.ListUser(appctx.GetConnectionToDatabase())
		if err != nil {
			panic(err)
		}
		// Handle home route
		//c.HTML(http.StatusOK, ".\\templates\\home_page.html", nil)
		//c.JSON(http.StatusOK, "home.tmpl", users)
		c.HTML(http.StatusOK, "home_page.tmpl", gin.H{"users": users})
	}
}
