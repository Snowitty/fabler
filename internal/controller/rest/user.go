package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/snowitty/fabler/internal/h"
	"github.com/snowitty/fabler/internal/model"
	"strconv"
)

func GetUser(c *gin.Context){
	var u model.User
	var err error

	var id = c.Params("id")

	u.ID, err = strconv.Atoi(id)
	if err != nil{
		c.JSON(200, h.Response{
			Status: 500,
			Message: err.Error(),
		})
		return
	}

	if user, err := u.Get(); err != nil{
		c.JSON(200, h.Response{
			Status: 404,
			Message: err.Error(),
		})
	}else {
		c.JSON(200, h.Response{
			Status: 200,
			Message: user,
		})
	}
	return
}


