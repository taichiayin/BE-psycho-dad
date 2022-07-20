package utils

import "psycho-dad/proto"

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RespPagin struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    int64       `json:"page"`
	Size    int64       `json:"size"`
	Total   int64       `json:"total"`
}

func RespSuccess(data interface{}) *Resp {
	res := &Resp{
		Code:    1,
		Message: "SUCCESSFUL",
		Data:    data,
	}
	return res
}

func PagingRespSuccess(data interface{}, p *proto.Paging) *RespPagin {
	res := &RespPagin{
		Code:    1,
		Message: "SUCCESSFUL",
		Data:    data,
		Page:    p.PageIndex,
		Size:    p.PageSize,
		Total:   p.AllCount,
	}
	return res
}

func RespError(errMessage string) *Resp {
	res := &Resp{
		Code:    0,
		Message: errMessage,
		Data:    nil,
	}
	return res
}
