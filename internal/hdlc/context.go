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
	ctx.currentState.Open()

	request := []byte{0xFA, 0xA5, 0x5F, 0xDC, 0x30, 0x81}
	ctx.io.Write(request)

	resp := make([]byte, 0)

	for {
		bytes := make([]byte, 128)

		n, err := ctx.io.Read(bytes)
		if err != nil {
			log.Fatal(err)
			break
		}

		if n == 0 {
			break
		}

		resp = append(resp, bytes[:n]...)
	}
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
