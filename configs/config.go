package configs

import "github.com/spf13/viper"

type Conf struct {
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	ViaCepApiUrl  string `mapstructure:"VIA_CEP_API_URL"`
	WeatherApiUrl string `mapstructure:"WEATHER_API_URL"`
}

func LoadConfig(path string) (*Conf, error) {

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var cfg *Conf
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
