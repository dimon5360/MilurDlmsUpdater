package ioservice

import (
	"app/main/internal/config"
	"fmt"
	"log"
	"sync"
	"time"
)

type TcpIO struct {
	Port string
	ctx *Context
}

func newTcpIO(device *config.Device, ctx *Context) *TcpIO {
	return &TcpIO{device.Port, ctx}
}

func (io *TcpIO) Init() {
	log.Printf("Open TPC port: %v\n", io.Port)

	// TODO: open TCP connection
}

func (io *TcpIO) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	io.ctx.Open()

	// long job imitation
	for i := 1; i <= 10; i++ {
		fmt.Println("Tcp port exchange ...")
		time.Sleep(time.Second)
		io.ctx.Exchange()
	}
    io.ctx.Close()
}