package helper

import "github.com/go-playground/validator/v10"

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(messageReq string, codeReq int, statusReq string, dataReq interface{}) response {
	meta := meta{
		Message: messageReq,
		Code:    codeReq,
		Status:  statusReq,
	}

	jsonResponse := response{
		Meta: meta,
		Data: dataReq,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
