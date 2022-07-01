package utils

type responseStruct struct {
	code    int
	message string
	data    interface{}
}

func HandleResponse(data) responseStruct {

	return responseStruct{code: 1, message: "Successful", data: &data}
}
