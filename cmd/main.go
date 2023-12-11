package cmd

import (
	"app/main/internal/config"
	ioservice "app/main/internal/io"
	"fmt"
	"sync"
)

func StartApp() {

	var wg sync.WaitGroup
	var appConfig config.App
	var devicesConfig []config.Device

	config.Parse("config/app.json", &appConfig)
	config.Parse("config/service.json", &devicesConfig)

	info := fmt.Sprintf("%s v.%s", appConfig.Appname, appConfig.Appversion)

	fmt.Println(info)

	for _, device := range devicesConfig {

		service := ioservice.Create(&device)
		ioservice.Init(service)

        wg.Add(1)
		go service.Run(&wg)
	}

	wg.Wait()
}
