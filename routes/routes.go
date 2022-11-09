package routes

import (
	"github.com/BINGSSJ/golang_todolist_project/api"
	"github.com/BINGSSJ/golang_todolist_project/midware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(midware.JWT())
		{
			authed.POST("task", api.CreateTask)
			authed.GET("task/:id", api.ShowTask)
			authed.GET("tasks", api.ListTask)
			authed.POST("task/:id", api.UpdateTask)
		}
	}
	return r
}
