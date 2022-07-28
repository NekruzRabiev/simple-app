package main

import "github.com/nekruzrabiev/simple-app/internal/app"

const configPath = "configs"

func main() {
	app.Run(configPath)
}
