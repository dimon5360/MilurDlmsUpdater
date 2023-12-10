package cmd

import (
	"app/main/internal/config"
	ioservice "app/main/internal/io"
	"fmt"
)

func StartApp() {

	var appConfig config.App
	var devicesConfig []config.Device

	config.Parse("config/app.json", &appConfig)
	config.Parse("config/service.json", &devicesConfig)

	info := fmt.Sprintf("%s v.%s", appConfig.Appname, appConfig.Appversion)

	fmt.Println(info)

	for _, device := range devicesConfig {
		ioservice.Init(ioservice.Create(&device))
	}

	ioservice.TestStates()
}
