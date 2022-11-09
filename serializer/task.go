package serializer

import "github.com/BINGSSJ/golang_todolist_project/dbmodel"

type Task struct {
	ID      uint   `json:"id" example:"1"`
	Title   string `json:"title" example:"Dinner"`
	Content string `json:"content" example: "we are going to eat dinner"`
	//View      uint64 `json:"view" example:"32"`  //浏览量
	Status    int   `json:"status" example:"0"` // 状态， 0未完成
	CreateAt  int64 `json:"create_at"`
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}

func BuildTask(item dbmodel.Task) Task {
	return Task{
		ID:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		CreateAt:  item.CreatedAt.Unix(),
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}

func BuildTasks(items []dbmodel.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
