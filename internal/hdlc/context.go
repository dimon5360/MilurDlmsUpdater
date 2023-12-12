package hdlc

import (
	state "app/main/internal/hdlc/states"
	"log"
)

type Context struct {
	currentState state.State

	openState     state.State
	idleState     state.State
	exchangeState state.State
	closedState   state.State

	io IO
}

func NewContext(io IO) *Context {

	openState := &state.OpenState{}
	idleState := &state.IdleState{}
	exchangeState := &state.ExchangeState{}
	closedState := &state.ClosedState{}

	ctx := &Context{
		currentState: closedState,

		openState:     openState,
		idleState:     idleState,
		exchangeState: exchangeState,
		closedState:   closedState,

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

	addr, err := queryPhysicalAddress(ctx.io)
	if err != nil {
		ctx.setState(ctx.closedState)
		return;
	}

	log.Printf("Got physical address %v\n", addr)

	// TODO: process address size and address itself

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
