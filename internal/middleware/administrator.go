package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snowitty/fabler/internal/model"
)

func Administartor() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, exist := c.Get("uid")
		if !exist {
			c.Redirect(http.StatusTemporaryRedirect, "/error?message=Not exist your uid")
			c.Abort()
			return
		}

		value, ok := ID.(int)
		if !ok {
			c.Redirect(http.StatusTemporaryRedirect, "/error?message=Your uid is not int")
			c.Abort()
			return
		}

		var u model.User
		u.ID = value

		user, err := u.Get()
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/error?message=User get error")
			c.Abort()
			return
		}

		if user.Type != 100 {
			c.Redirect(http.StatusTemporaryRedirect, "/error?message=You are not Administartor")
			c.Abort()
			return
		}

		c.Next()
	}
}
