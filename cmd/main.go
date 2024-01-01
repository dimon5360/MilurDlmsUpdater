package cmd

import (
	"app/main/internal/config"
	ioservice "app/main/internal/io"
	"fmt"
	"log"
	"sync"
)

type app struct {
	wg sync.WaitGroup
	config config.Info
	devicesConfig []config.Device
}

func NewApp(path string) *app {

	a := app{}

	config.Parse(path, &a.config)

	fmt.Printf("%s v.%s\n", a.config.Appname, a.config.Appversion)
	fmt.Printf("Build datetime: %s\n\n", a.config.Builddate)
	return &a
}

func (a *app) Config(path string) *app {

	config.Parse(path, &a.devicesConfig)
	return a
}

func (a *app) Run() {

	for _, device := range a.devicesConfig {

		service := ioservice.Create(&device)
		err := ioservice.Init(service)

		if err != nil {
			log.Println("Service creation failed")
			continue
		}

        a.wg.Add(1)
		go service.Run(&a.wg)
	}

	a.wg.Wait()
}
