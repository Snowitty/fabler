package rest

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/snowitty/fabler/internal/h"
	"github.com/snowitty/fabler/internal/model"
)

func GetUser(c *gin.Context) {
	var u model.User
	var err error

	var id = c.Param("id")

	u.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if user, err := u.Get(); err != nil {
		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})
	} else {
		c.JSON(200, h.Response{
			Status:  200,
			Message: user,
		})
	}
	return
}

func GetUsers(c *gin.Context) {

	var u model.User
	var limmit int
	var offset int
	var err error

	limmit, err = strconv.Atoi(c.Query("limmit"))
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	offset, err = strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if users, err := u.GetList(limmit, offset); err != nil {
		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})
	} else {
		c.JSON(200, h.Response{
			Status:  200,
			Message: users,
		})
	}

	return
}

func GetUsersCounts(c *gin.Context) {
	var u model.User

	if counts, err := u.GetCounts(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})
	} else {

		c.JSON(200, h.Response{
			Status:  200,
			Message: counts,
		})
	}
	return
}

func UpdateUser(c *gin.Context) {
	var u model.User
	var ra int64
	var err error

	var id = c.Param("id")

	u.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&u); err != nil {
		c.JSON(200, h.Response{Status: 500, Message: err.Error()})
		return
	}

	if ra, err = u.Update(); err != nil {
		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})
	} else {
		c.JSON(200, h.Response{
			Status:  200,
			Message: ra,
		})
	}
	return
}

func DeleteUser(c *gin.Context) {
	var u model.User

	var id = c.Param("id")

	var ra int64
	var err error

	u.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if ra, err = u.Delete(); err != nil {
		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})
	} else {
		c.JSON(200, h.Response{
			Status:  200,
			Message: ra,
		})
	}

	return
}
