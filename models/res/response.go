package res

import (
	"Backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success = 0
	Error   = 7
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type ListResponse struct {
	Count int64 `json:"count"`
	List  any   `json:"list"`
}

func Result(code int, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data interface{}, msg string, ctx *gin.Context) {
	Result(Success, data, msg, ctx)
}

func OkWith(ctx *gin.Context) {
	Result(Success, map[string]interface{}{}, "成功", ctx)
}

func OkWithData(data interface{}, ctx *gin.Context) {
	Result(Success, data, "成功", ctx)
}

func OkWithMessage(msg string, ctx *gin.Context) {
	Result(Success, map[string]interface{}{}, msg, ctx)
}

func OkWithList(list any, count int64, ctx *gin.Context) {
	OkWithData(ListResponse{
		Count: count,
		List:  list,
	}, ctx)
}

func Fail(data interface{}, msg string, ctx *gin.Context) {
	Result(Error, data, msg, ctx)
}

func FailWithMessage(msg string, ctx *gin.Context) {
	Result(Error, map[string]interface{}{}, msg, ctx)
}

func FailWithCode(code ErrorCode, ctx *gin.Context) {
	msg, err := ErrorMap[ErrorCode(code)]
	if err {
		Result(int(code), map[string]interface{}{}, msg, ctx)
	}
	Result(Error, map[string]interface{}{}, msg, ctx)
}

func FailWithError(err error, obj any, ctx *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, ctx)
}
