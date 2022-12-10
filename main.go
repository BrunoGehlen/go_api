package main

import (
	"github.com/BrunoGehlen/go_api/model/app"
	"github.com/BrunoGehlen/go_api/model/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
