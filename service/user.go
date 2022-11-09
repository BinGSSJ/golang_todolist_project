package service

import (
	"github.com/BINGSSJ/golang_todolist_project/dbmodel"
	"github.com/BINGSSJ/golang_todolist_project/pkg/utils"
	"github.com/BINGSSJ/golang_todolist_project/serializer"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	PassWord string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user dbmodel.User
	var count int
	dbmodel.DB.Model(&dbmodel.User{}).Where("user_name=?",
		service.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "user already exists",
		}
	}
	user.UserName = service.UserName
	if err := user.SetPwd(service.PassWord); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	// 创建新用户
	if err := dbmodel.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "database error",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "Register Success",
	}
}

func (service *UserService) Login() serializer.Response {
	var user dbmodel.User
	// 查询用户是否存在
	if err := dbmodel.DB.Where("user_name=?",
		service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "user not exits!",
			}
		} else {
			return serializer.Response{
				Status: 500,
				Msg:    "database ERROR!",
			}
		}
	}
	if user.CheckPwd(service.PassWord) == false {
		return serializer.Response{
			Status: 400,
			Msg:    "Password ERROR!",
		}
	}
	// 验证成功，发送token
	token, err := utils.GenerateToken(user.ID, service.UserName)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "token generate ERROR",
		}
	}

	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: "Login Success",
	}
}
