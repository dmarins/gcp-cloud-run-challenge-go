package main

import (
	"fmt"

	"github.com/dmarins/gcp-cloud-run-challenge-go/configs"
	"github.com/dmarins/gcp-cloud-run-challenge-go/internal/infrastructure/web/server"
)

func main() {
	// Envs
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	getWeatherByZipcodeUseCase := NewGetWeatherByZipcodeUseCase()
	weatherHttpHandler := NewWeatherHttpHandler(*getWeatherByZipcodeUseCase)

	// Http Server
	webserver := server.NewWebServer(configs.WebServerPort)

	webserver.AddHandler("get", "/weather/{zipcode}", weatherHttpHandler.GetWeatherInfo)

	fmt.Println("Starting HTTP server on port", configs.WebServerPort)

	webserver.Start()
}
