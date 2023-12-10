package ioservice

import (
	state "app/main/internal/io/states"
	"log"
)

type Context struct {
    currentState 	state.State

    openState 		state.State
    idleState 		state.State
    exchangeState 	state.State
    closedState 	state.State
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

func (ctx *Context) Open() {
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

func TestStates() {
	ctx := newContext()

    ctx.Open()
    ctx.Exchange()
    ctx.Exchange()
    ctx.Exchange()
    ctx.Close()
}
