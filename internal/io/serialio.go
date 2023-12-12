package ioservice

import (
	"app/main/internal/config"
	"log"
	"slices"
	"sync"

	"go.bug.st/serial"
)

type SerialIO struct {
	Port 	string
	Speed 	int

	port 	serial.Port
	ctx 	*Context
}

func newSerialIO(device *config.Device) *SerialIO {
	return &SerialIO{device.Port, device.Speed, nil, &Context{}}
}

func (io *SerialIO) Init() error {

	ports, err := serial.GetPortsList()
	if err != nil {
		log.Println(err)
		return err
	}

	if len(ports) == 0 {
		log.Println("No serial ports found!")
		return err
	}

	if slices.Contains(ports, io.Port) {
		log.Printf("No serial %s found!", io.Port)
		return err
	}

	log.Printf("Found port: %v\n", io.Port)

	port, err := serial.Open(io.Port, &serial.Mode{BaudRate: io.Speed})
	if err != nil {
		log.Println(err)
		return err
	}

	io.port = port
	return nil
}

func (io *SerialIO) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	io.ctx.Init(io.port)
	io.ctx.Open()
    io.ctx.Close()

	io.port.Close()
}
