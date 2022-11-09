package service

import (
	"github.com/BINGSSJ/golang_todolist_project/dbmodel"
	"github.com/BINGSSJ/golang_todolist_project/serializer"
	"time"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` // 0 unfinished, 1 finished
}

type ShowTaskService struct {
}
type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` // 0 unfinished, 1 finished
}

func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user dbmodel.User
	var code int
	dbmodel.DB.First(&user, id)
	task := dbmodel.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
	}
	err := dbmodel.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "Create task Failed",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "Create task Success",
	}
}

func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task dbmodel.Task
	code := 200
	err := dbmodel.DB.First(&task, tid).Error
	if err != nil {
		code = 400
		return serializer.Response{
			Status: code,
			Msg:    "query Fail",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
	}
}

func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []dbmodel.Task
	count := 0
	if service.PageNum == 0 {
		service.PageSize = 15
	}
	dbmodel.DB.Model(&dbmodel.Task{}).Preload("User").Where(
		"uid=?", uid).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

func (service *UpdateTaskService) Update(tid string) serializer.Response {
	var task dbmodel.Task
	dbmodel.DB.First(&task, tid)
	task.Content = service.Content
	task.Status = service.Status
	task.Title = service.Title
	err := dbmodel.DB.Save(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "update Fail",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
		Msg:    "update Success",
	}

}
