package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResult struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(status int, code int, data any, msg string, c *gin.Context) {
	c.JSON(status, JsonResult{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Success(data any, c *gin.Context) {
	Result(http.StatusOK, http.StatusOK, data, StatusMsg[http.StatusOK], c)
}

func Failed(code int, msg string, c *gin.Context) {
	Result(http.StatusOK, code, map[string]string{}, msg, c)
}

func FailedWithCode(code int, c *gin.Context) {
	if msg, ok := StatusMsg[code]; ok {
		Failed(code, msg, c)
	} else {
		Failed(code, "", c)
	}
}

func FailedWithMsg(msg string, c *gin.Context) {
	Failed(0, msg, c)
}
