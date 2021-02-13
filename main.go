package main

import (
	"github.com/dscontrak/go-gorm-mux/api/app"
	"github.com/dscontrak/go-gorm-mux/api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
