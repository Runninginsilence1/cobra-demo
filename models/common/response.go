package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	FAILED  = 4001
	SUCCESS = 0
)

type Response struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data,omitempty"`
	Msg       string      `json:"msg"`
	Timestamp string      `json:"ts"`
}

func result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
		time.Now().Format("2006-01-02 15:04:05"),
	})
}

func OK(c *gin.Context) {
	result(SUCCESS, nil, "操作成功", c)
}

func OKWithDetail(msg string, data any, c *gin.Context) {
	result(SUCCESS, data, msg, c)
}

func Failed(c *gin.Context) {
	result(FAILED, nil, "操作失败", c)
}

func FailedWithWrongReq(c *gin.Context) {
	result(FAILED, nil, "JSON格式不正确", c)
}

func FailedWithDetail(msg string, data any, c *gin.Context) {
	result(FAILED, data, msg, c)
}
