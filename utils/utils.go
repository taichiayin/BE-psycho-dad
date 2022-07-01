package utils

import (
	"fmt"
)

type ResponseStruct struct {
	code    int
	message string
	data    interface{}
}

func HandleResponse(data interface{}) ResponseStruct {
	fmt.Println(data)

	return ResponseStruct{code: 1, message: "Successful", data: &data}
}
