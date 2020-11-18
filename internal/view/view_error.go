package view

import (
	"github.com/gin-gonic/gin"
	"github.com/snowitty/fabler/conf"
	"net/http"
)

func Error(c *gin.Context){
	var csrdata map[string]interface{}

	template := "error.html"
	data := map[string]interface{}{
		"lang": conf.Config().Lang,
		"csrdata": csrdata,
	}
	c.HTML(http.StatusOK, template, data)

	return
}
