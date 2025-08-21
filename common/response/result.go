package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zsljava/gokit/util"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"trace_id"`
}

func Success(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	resp := Response{Code: 0, Message: "ok", Data: data, TraceId: util.GetTraceId(ctx)}
	ctx.JSON(http.StatusOK, resp)
}

func Error(ctx *gin.Context, httpCode int, err error, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := Response{Code: errorCodeMap[err], Message: err.Error(), Data: data, TraceId: util.GetTraceId(ctx)}
	if _, ok := errorCodeMap[err]; !ok {
		resp = Response{Code: 500, Message: "unknown common", Data: data}
	}
	ctx.JSON(httpCode, resp)
}

type RespError struct {
	Code    int
	Message string
}

var errorCodeMap = map[error]int{}

func NewError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}
func (e RespError) Error() string {
	return e.Message
}
