package ioservice

import "app/main/internal/hdlc"

type Context struct {
	hdlc *hdlc.Context
}

func (ctx *Context) Init(io hdlc.IO) {
	ctx.hdlc = hdlc.NewContext(io)
}

func (ctx *Context) Open() {
	ctx.hdlc.Open()
}

func (ctx *Context) Exchange() {

}

func (ctx *Context) Close() {

}
