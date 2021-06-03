package middleware

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SessionValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := sessions.Default(c).Get("user")
		if user == nil {
			c.Redirect(http.StatusFound, fmt.Sprintf("login?rand=%d", time.Now().UnixNano()))
		}
	}
}
