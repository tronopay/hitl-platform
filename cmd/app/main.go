package main

import (
	_ "tronopay/docs"
	"tronopay/internal/app"
)

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}
