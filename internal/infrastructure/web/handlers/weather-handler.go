package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dmarins/gcp-cloud-run-challenge-go/internal/usecase"
	"github.com/go-chi/chi"
)

type WeatherHttpHandler struct {
	GetWeatherByZipcodeUseCase usecase.GetWeatherByZipcodeUseCase
}

func NewWeatherHttpHandler(getWeatherByZipcodeUseCase usecase.GetWeatherByZipcodeUseCase) *WeatherHttpHandler {
	return &WeatherHttpHandler{
		GetWeatherByZipcodeUseCase: getWeatherByZipcodeUseCase,
	}
}

func (handler *WeatherHttpHandler) GetWeatherInfo(w http.ResponseWriter, r *http.Request) {

	zipcode := chi.URLParam(r, "zipcode")
	if zipcode == "" {
		http.Error(w, "invalid zipcode parameter.", http.StatusBadRequest)
		return
	}

	dto := usecase.WeatherInputDTO{
		Zipcode: zipcode,
	}

	output, err := handler.GetWeatherByZipcodeUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
