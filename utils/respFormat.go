package utils

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RespSuccess(data interface{}) *Resp {
	res := &Resp{
		Code:    1,
		Message: "SUCCESSFUL",
		Data:    data,
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
