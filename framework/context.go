package framework

import (
	"context"
	"net/http"
	"sync"
)

// 自定义context
type Context struct {
	request  *http.Request
	response http.ResponseWriter
	ctx      context.Context
	handler  ControllerHandler

	// 是否超时标记位
	hasTimeout bool
	// 写保护机制
	writerMux *sync.Mutex
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		request:   r,
		response:  w,
		ctx:       r.Context(),
		writerMux: &sync.Mutex{},
	}
}

func (ctx Context) WriterMutex() *sync.Mutex {
	return ctx.writerMux
}

func (ctx Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx Context) GetResponse() http.ResponseWriter {
	return ctx.response
}

func (ctx Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}
