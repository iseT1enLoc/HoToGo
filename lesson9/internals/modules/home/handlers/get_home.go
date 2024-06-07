package homepagehandlers

import (
	appcontext1 "lesson9/internals/common/appcontext"
	"lesson9/internals/templates"
	"lesson9/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHomeBase(appctx appcontext1.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("token"); err == nil && cookie != "" {
			utils.Render(c, http.StatusOK, templates.Index(true))
		} else {
			utils.Render(c, http.StatusOK, templates.Index(false))
		}
	}
}
