package main

import (
	"app/main/cmd"
)

const (
	app_config = "config/app.json"
	service_config = "config/service.json"
)

func main() {
	app := cmd.NewApp(app_config)
	app.Config(service_config).Run()
}
