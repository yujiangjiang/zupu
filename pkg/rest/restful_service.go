package rest

import "github.com/gin-gonic/gin"

type RS struct {
	e *gin.Engine
}

func NewRS(e *gin.Engine) RS {
	return RS{
		e: e,
	}
}

func (rs *RS) Group(relativePath string) *gin.RouterGroup {
	return rs.e.Group(relativePath)
}

type Handler func(ctx Context)

func (rs *RS) Request(method string, group *gin.RouterGroup, uri string, handler Handler) {
	h := func(context *gin.Context) {
		handler(Context{context})
	}
	group.Handle(method, uri, h)
}

func (rs *RS) GET(group *gin.RouterGroup, uri string, handler Handler) {
	rs.Request("GET", group, uri, handler)
}

func (rs *RS) POST(group *gin.RouterGroup, uri string, handler Handler) {
	rs.Request("POST", group, uri, handler)
}

func (rs *RS) PUT(group *gin.RouterGroup, uri string, handler Handler) {
	rs.Request("PUT", group, uri, handler)
}

func (rs *RS) DELETE(group *gin.RouterGroup, uri string, handler Handler) {
	rs.Request("DELETE", group, uri, handler)
}

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func NewResponse() Response {
	return Response{
		Code:   200,
		Msg:    "success",
		Result: nil,
	}
}

type Context struct {
	Ctx *gin.Context
}

func (ctx *Context) Success() {
	ctx.SuccessWithResult(nil)
}

func (ctx *Context) SuccessWithResult(result interface{}) {
	ctx.Json(200, "success", result)
}

func (ctx *Context) SystemError() {
	ctx.Json(500, "system error", nil)
}

func (ctx *Context) ValidateError(msg string) {
	ctx.Json(400, msg, nil)
}

func (ctx *Context) Json(code int, msg string, result interface{}) {
	ctx.Ctx.JSON(200, Response{
		Code:   code,
		Msg:    msg,
		Result: result,
	})
}
