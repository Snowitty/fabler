package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snowitty/fabler/internal/constant"
	"github.com/snowitty/fabler/internal/service/authorization"
)

func Authorizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ss, err := c.Cookie(constant.SSKEY)
		if err != nil {
			log.Print("Authorizer error: ", err)
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		ID, ProfileID, err := authorization.Parse(ss)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		c.Set("uid", ID)
		c.Set("pid", ProfileID)

		c.Next()

		log.Print("UID: ", ID)
		log.Print("PID: ", ProfileID)

	}
}
