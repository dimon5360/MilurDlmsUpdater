package main

import (
	"app/main/cmd"
)

func main() {

	cmd.App().
		Config().
		Run()
}
