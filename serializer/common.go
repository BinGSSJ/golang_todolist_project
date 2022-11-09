package serializer

type Response struct { // 序列化基本
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {//带总数的struct
	Item interface{} `json:"item"`
	Total uint `json:"total"`
}

func BuildListResponse(items interface{}, total uint) Response{
	return Response{
		Status: 200,
		Data: DataList{
			Item: items,
			Total: total,
		},
	}
}
