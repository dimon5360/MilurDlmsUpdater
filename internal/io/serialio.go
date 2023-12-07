package ioservice

import "fmt"

type SerialIO struct {
	Port string
}

func (io *SerialIO) Init() {
	fmt.Println("Init Serial IO connection")

	// TODO: init serial connection
}
