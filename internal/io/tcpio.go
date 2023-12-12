package ioservice

import (
	"app/main/internal/config"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type TcpIO struct {
	Port string

	port *net.TCPConn
	ctx  *Context
}

func newTcpIO(device *config.Device, ctx *Context) *TcpIO {
	return &TcpIO{device.Port, nil, ctx}
}

func (io *TcpIO) Init() {
	log.Printf("Open TPC port: %v\n", io.Port)

    tcpAddr, err := net.ResolveTCPAddr("tcp", io.Port)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        log.Fatalf("Dial failed: %v\n", err.Error())
        os.Exit(1)
    }

	io.port = conn
    conn.Close()
}

func (io *TcpIO) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	io.ctx.Open(io.port)

	// long job imitation
	for i := 1; i <= 10; i++ {
		fmt.Println("Tcp port exchange ...")

		io.port.Write([]byte("10,20,30\n\r"))
		buf := make([]byte, 128)
		io.port.Read(buf)

		time.Sleep(time.Second)
		io.ctx.Exchange()
	}
    io.ctx.Close()
}