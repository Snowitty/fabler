package view

import (
	"github.com/gin-gonic/gin"
	"github.com/snowitty/fabler/conf"
	"net/http"
)

func Login(c *gin.Context){

	var csrdata map[string]interface{}

	template := "login.html"
	data := map[string]interface{}{
		"lang": conf.Config().Lang,
		"csrdata": csrdata,
	}

	c.HTML(http.StatusOK, template, data)

	return
}
