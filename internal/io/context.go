package ioservice

import (
	state "app/main/internal/io/states"
	"log"
)

type Connection interface {
	Write([]byte) (int, error)
	Read([]byte) (int, error)
}

type Context struct {
    currentState 	state.State

    openState 		state.State
    idleState 		state.State
    exchangeState 	state.State
    closedState 	state.State

	conn 			Connection
}

func newContext() *Context {

    openState := &state.OpenState{}
    idleState := &state.IdleState{}
    exchangeState := &state.ExchangeState{}
    closedState := &state.ClosedState{}

    ctx := &Context{
        currentState: closedState,

		openState: openState,
        idleState: idleState,
        exchangeState: exchangeState,
        closedState: closedState,
    }
	return ctx
}

func (ctx *Context) setState(state state.State) {
    ctx.currentState = state
	log.Println("New state:", ctx.currentState.GetDescription())
}

func (ctx *Context) Open(conn Connection) {
	ctx.conn = conn

	ctx.setState(ctx.openState)
    ctx.currentState.Open()
	ctx.setState(ctx.idleState)
}

func (ctx *Context) Exchange() {
	ctx.setState(ctx.exchangeState)
    ctx.currentState.Exchange()
	ctx.setState(ctx.idleState)
}

func (ctx *Context) Close() {
	ctx.setState(ctx.closedState)
    ctx.currentState.Close()
}
