package cmd

import (
	"app/main/internal/config"
	ioservice "app/main/internal/io"
	"fmt"
	"log"
	"sync"
)

func StartApp() {

	var wg sync.WaitGroup
	var appConfig config.App
	var devicesConfig []config.Device

	config.Parse("config/app.json", &appConfig)
	config.Parse("config/service.json", &devicesConfig)

	fmt.Printf("%s v.%s\n", appConfig.Appname, appConfig.Appversion)
	fmt.Printf("Build datetime: %s\n\n", appConfig.Builddate)

	for _, device := range devicesConfig {

		service := ioservice.Create(&device)
		err := ioservice.Init(service)

		if err != nil {
			log.Println("Service creation failed")
			continue
		}

        wg.Add(1)
		go service.Run(&wg)
	}

	wg.Wait()
}
