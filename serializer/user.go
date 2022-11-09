package serializer

import "github.com/BINGSSJ/golang_todolist_project/dbmodel"

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	UserName string `json:"user_name" form:"user_name" example:"bing"`
	Status   string `json:"status" form:"status"`
	CreateAt int64  `json:"create_at" form:"create_at"`
}

func BuildUser(user dbmodel.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
