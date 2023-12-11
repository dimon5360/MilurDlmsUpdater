package ioservice

import (
	"app/main/internal/config"
	"fmt"
	"sync"
)

type IO struct {
	Serial *SerialIO
	Tcp    *TcpIO
}

type Service interface {
	Init()
	Run(wg *sync.WaitGroup)
}

func Create(device *config.Device) Service {

	switch device.Type {
	case "Serial":
		{
			fmt.Println("Create Serial IO service")
			return newSerialIO(device, newContext())
		}
	case "TCP":
		{
			fmt.Println("Create TCP IO service")
			return newTcpIO(device, newContext())
		}
	default:
		{
			panic("Undefined service")
		}
	}
}

func Init(s Service) {
	s.Init()
}

func Run(s Service, wg *sync.WaitGroup) {
	s.Run(wg)
}