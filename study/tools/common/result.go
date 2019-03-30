package common

import "net/http"

const (
	ServerError         = 500      //服务端异常
	StatusParamError    = 10904001 //请求参数错误
	StatusDatabaseError = 10904002 // 数据库异常
)

type Result struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Metadata Metadata    `json:"metadata"`
	Data     interface{} `json:"data"`
}

type Metadata struct {
	TotalRows   int `json:"totalRows"`
	CurrentPage int `json:"currentPage"`
	PageRows    int `json:"pageRows"`
}

func Success(data interface{}) *Result {
	return &Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    data,
	}
}

func SuccessMsg(data interface{}, msg string) *Result {
	return &Result{
		Status:  http.StatusOK,
		Message: msg,
		Data:    data,
	}
}

func FailWithData(code int, data interface{}) *Result {
	return &Result{
		Status:  code,
		Message: "FAIL",
		Data:    data,
	}
}

func FailWithMsg(code int, message string) *Result {
	return &Result{
		Status:  code,
		Message: message,
		Data:    nil,
	}
}

func IsSuccess(result *Result) bool {
	return result.Status == http.StatusOK
}
