//go:build wireinject
// +build wireinject

package main

import (
	"github.com/dmarins/gcp-cloud-run-challenge-go/internal/infrastructure/web/handlers"
	"github.com/dmarins/gcp-cloud-run-challenge-go/internal/usecase"
	"github.com/google/wire"
)

func NewGetWeatherByZipcodeUseCase() *usecase.GetWeatherByZipcodeUseCase {
	wire.Build(
		usecase.NewGetWeatherByZipcodeUseCase,
	)

	return &usecase.GetWeatherByZipcodeUseCase{}
}

func NewWeatherHttpHandler(getWeatherByZipcodeUseCase usecase.GetWeatherByZipcodeUseCase) *handlers.WeatherHttpHandler {
	wire.Build(
		handlers.NewWeatherHttpHandler,
	)

	return &handlers.WeatherHttpHandler{}
}
