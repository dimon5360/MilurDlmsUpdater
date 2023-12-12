package ioservice

import (
	"app/main/internal/config"
	"fmt"
	"log"
	"slices"
	"sync"
	"time"

	"go.bug.st/serial"
)

type SerialIO struct {
	Port  string
	Speed int

	port serial.Port
	ctx  *Context
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

	log.Println("Free serial ports:")
	for _, port := range ports {
		log.Printf("\t%v\n", port)
	}

	if slices.Contains(ports, io.Port) {
		log.Printf("No serial %s found!", io.Port)
		return fmt.Errorf("%s: %s", "Serial port found", io.Port)
	}

	log.Printf("Found port: %v\n", io.Port)

	port, err := serial.Open(io.Port, &serial.Mode{BaudRate: io.Speed})
	if err != nil {
		log.Println(err)
		return err
	}

	io.port = port

	return io.port.SetReadTimeout(time.Duration(100) * time.Millisecond)
}

func (io *SerialIO) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	io.ctx.Init(io.port)
	io.ctx.Open()
	io.ctx.Close()

	io.port.Close()
}
