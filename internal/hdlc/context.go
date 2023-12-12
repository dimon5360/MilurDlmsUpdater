package hdlc

import (
	state "app/main/internal/hdlc/states"
	"log"
)

type Context struct {
    currentState 	state.State

    openState 		state.State
    idleState 		state.State
    exchangeState 	state.State
    closedState 	state.State

	io IO
}

func NewContext(io IO) *Context {

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

		io: io,
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

	ctx.io.Write([]byte("10,20,30\n\r"))
	resp := make([]byte, 128)
	ctx.io.Read(resp)

	log.Printf("Got bytes %v\n", resp)

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