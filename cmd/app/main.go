package main

import "github.com/nekruzrabiev/simple-app/internal/app"

const configPath = "configs"

// @title Simple App API
// @version 1.0
// @description API Server for simple-app

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.Run(configPath)
}
