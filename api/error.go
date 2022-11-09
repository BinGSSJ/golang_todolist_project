package api

import (
	"encoding/json"
	"fmt"
	"github.com/BINGSSJ/golang_todolist_project/serializer"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 401,
			Msg:    "json ERROR",
			Error:  fmt.Sprint(err),
		}
	}
	return serializer.Response{
		Status: 401,
		Msg:    "params ERROR",
		Error:  fmt.Sprint(err),
	}
}
