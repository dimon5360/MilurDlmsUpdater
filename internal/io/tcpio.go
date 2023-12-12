package ioservice

import (
	"app/main/internal/config"
	"log"
	"net"
	"sync"
)

type TcpIO struct {
	Port string

	port *net.TCPConn
	ctx  *Context
}

func newTcpIO(device *config.Device) *TcpIO {
	return &TcpIO{device.Port, nil, &Context{}}
}

func (io *TcpIO) Init() error {
	log.Printf("Open TPC port: %v\n", io.Port)

	tcpAddr, err := net.ResolveTCPAddr("tcp", io.Port)
	if err != nil {
		log.Println("ResolveTCPAddr failed:", err.Error())
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("Dial failed: %v\n", err.Error())
		return err
	}

	io.port = conn
	return nil
}

func (io *TcpIO) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	io.ctx.Init(io.port)
	io.ctx.Open()
	io.ctx.Close()

	io.port.Close()
}
