package main

import (
	"github.com/BINGSSJ/golang_todolist_project/conf"
	"github.com/BINGSSJ/golang_todolist_project/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
