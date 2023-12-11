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
	Port 	string
	Speed 	int
	ctx 	*Context
}

func newSerialIO(device *config.Device, ctx *Context) *SerialIO {
	return &SerialIO{device.Port, device.Speed, ctx}
}

func (io *SerialIO) Init() {

	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}

	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	if slices.Contains(ports, io.Port) {
		log.Fatalf("No serial %s found!", io.Port)
	}

	log.Printf("Found port: %v\n", io.Port)

	port, err := serial.Open(io.Port, &serial.Mode{BaudRate: io.Speed})
	if err != nil {
		log.Fatal(err)
	}

	// test string
	n, err := port.Write([]byte("10,20,30\n\r"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sent %d bytes\n", n)
}

func (io *SerialIO) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	io.ctx.Open()

	// long job imitation
	for i := 1; i <= 3; i++ {
		fmt.Println("Serial port exchange ...")
		time.Sleep(time.Second)
		io.ctx.Exchange()
	}
    io.ctx.Close()
}
