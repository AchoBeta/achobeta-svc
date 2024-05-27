package web

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonMsgResponse struct {
	Ctx *gin.Context
}

func JsonMsgResult() map[string]interface{} {
	return make(map[string]interface{})
}

const SUCCESS_CODE = 200
const SUCCESS_MSG = "success"
const ERROR_MSG = "error"
const CODE = "code"
const MESSAGE = "message"
const DATA = "data"
const TRACEID = "traceId"

func NewResponse(c *gin.Context) *JsonMsgResponse {
	return &JsonMsgResponse{Ctx: c}
}

func (r *JsonMsgResponse) Success(data interface{}) {
	res := JsonMsgResult()
	res[CODE] = constant.SUCCESS.Code
	res[MESSAGE] = constant.SUCCESS.Msg
	if data != nil {
		res[DATA] = data
	}
	r.Ctx.JSON(http.StatusOK, res)
}

func (r *JsonMsgResponse) Error() {
	r.error(constant.COMMON_FAIL.Code, constant.COMMON_FAIL.Msg)
}

func (r *JsonMsgResponse) ErrorCode(mc constant.MsgCode) {
	r.error(mc.Code, mc.Msg)
}

func (r *JsonMsgResponse) ErrorMsg(mc constant.MsgCode, message string) {
	r.error(mc.Code, message)
}

func (r *JsonMsgResponse) ErrorTrace(mc constant.MsgCode, message string, traceId string) {
	r.errorCtx(mc.Code, message, traceId)
}

func (r *JsonMsgResponse) errorCtx(code int, message string, traceId string) {
	if message == "" {
		message = constant.COMMON_FAIL.Msg
	}
	res := JsonMsgResult()
	res[CODE] = code
	res[MESSAGE] = message
	res[TRACEID] = traceId
	r.Ctx.JSON(http.StatusOK, res)
}

func (r *JsonMsgResponse) error(code int, message string) {
	if message == "" {
		message = constant.COMMON_FAIL.Msg
	}
	res := JsonMsgResult()
	res[CODE] = code
	res[MESSAGE] = message
	r.Ctx.JSON(http.StatusOK, res)
}
