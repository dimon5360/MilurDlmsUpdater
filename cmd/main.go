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

func App() *app {
	return &app{}
}

func (a *app) Config() *app {

	config.Parse("config/app.json", &a.config)
	config.Parse("config/service.json", &a.devicesConfig)

	fmt.Printf("%s v.%s\n", a.config.Appname, a.config.Appversion)
	fmt.Printf("Build datetime: %s\n\n", a.config.Builddate)

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
