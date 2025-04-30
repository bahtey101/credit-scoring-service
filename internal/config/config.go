package config

type Config struct {
	Host         string `env:"HOST,required"`
	Port         string `env:"PORT,required"`
	GoogleAPIKey string `env:"GOOGLE_API_KEY,required"`

	MLHost string `env:"ML_MODEL_HOST,required"`
	MLPort string `env:"ML_MODEL_PORT,required"`

	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}
