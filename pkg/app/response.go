package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ljinf/template_project_v2/pkg/errcode"
	"github.com/ljinf/template_project_v2/pkg/log"
	"net/http"
)

type Response struct {
	ctx       *gin.Context
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	RequestId string      `json:"request_id"`
	Data      interface{} `json:"data,omitempty"`
}

func HandleSuccess(c *gin.Context, data interface{}) {
	r := &Response{
		ctx:  c,
		Code: errcode.Success.Code(),
		Msg:  errcode.Success.Msg(),
		Data: data,
	}

	requestId := ""
	if _, exists := r.ctx.Get("traceid"); exists {
		val, _ := r.ctx.Get("traceid")
		requestId = val.(string)
	}
	r.RequestId = requestId
	r.ctx.JSON(http.StatusOK, r)
}

func HandleSuccessOk(c *gin.Context) {
	HandleSuccess(c, "")
}

func HandleError(c *gin.Context, err *errcode.AppError) {
	r := &Response{
		ctx:  c,
		Code: err.Code(),
		Msg:  err.Msg(),
	}
	requestId := ""
	if _, exists := r.ctx.Get("traceid"); exists {
		val, _ := r.ctx.Get("traceid")
		requestId = val.(string)
	}
	r.RequestId = requestId
	// 兜底记一条响应错误, 项目自定义的AppError中有错误链条, 方便出错后排查问题
	log.Error(r.ctx, "api_response_error", "err", err.Error())
	r.ctx.JSON(err.HttpStatusCode(), r)
}
