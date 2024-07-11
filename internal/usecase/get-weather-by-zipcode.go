package usecase

type (
	WeatherInputDTO struct {
		Zipcode string `json:"zipcode"`
	}

	ZipcodeOutputDTO struct {
		Location string `json:"localidade"`
	}

	WeatherOutputDTO struct {
		Celsius float64 `json:"celsius"`
	}

	GetWeatherByZipcodeUseCase struct{}
)

func NewGetWeatherByZipcodeUseCase() *GetWeatherByZipcodeUseCase {
	return &GetWeatherByZipcodeUseCase{}
}

func (usecase *GetWeatherByZipcodeUseCase) Execute(input WeatherInputDTO) (WeatherOutputDTO, error) {
	return WeatherOutputDTO{
		Celsius: 22.5,
	}, nil
}
