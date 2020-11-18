package view

import (
	"github.com/gin-gonic/gin"
	"github.com/snowitty/fabler/conf"
	"net/http"
)

func Admin(c *gin.Context){
	var csrdata map[string]interface{}

	template := "admin.html"
	data := map[string]interface{}{
		"lang": conf.Config().Lang,
		"csrdata": csrdata,
	}

	c.HTML(http.StatusOK, template, data)

	return
}
