package ioservice

import (
	"app/main/internal/config"
	"fmt"
)

type IO struct {
	Serial *SerialIO
	Tcp    *TcpIO
}

type Service interface {
	Init()
}

var dict map[string]IO

func Create(device *config.Device) (Service, *config.Device) {

	if dict == nil {
		dict = make(map[string]IO)
	}

	switch device.Type {
	case "Serial":
		{
			fmt.Println("Create Serial IO service")
			dict[device.Name] = IO{Serial: &SerialIO{device.Port}, Tcp: nil}
			return dict[device.Name].Serial, device
		}
	case "TCP":
		{
			fmt.Println("Create TCP IO service")
			dict[device.Name] = IO{Serial: nil, Tcp: &TcpIO{device.Port}}
			return dict[device.Name].Tcp, device
		}
	default:
		{
			panic("Undefined service")
		}
	}
}

func Init(s Service, device *config.Device) {
	fmt.Println("Init device", device.Name)
	s.Init()
}

func GetIO(device *config.Device) IO {
	return dict[device.Name]
}
