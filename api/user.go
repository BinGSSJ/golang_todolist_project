package api

import (
	"github.com/BINGSSJ/golang_todolist_project/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	// 绑定服务对象
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(res.Status, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
