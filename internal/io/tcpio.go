package ioservice

import "fmt"

type TcpIO struct {
	Port string
}

func (io *TcpIO) Init() {
	fmt.Println("Init TCP IO connection")

	// TODO: init tcp connection
}
