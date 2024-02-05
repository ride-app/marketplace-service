package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Production           bool   `env:"PRODUCTION" env-description:"dev or prod" env-default:"true"`
	LogDebug             bool   `env:"LOG_DEBUG" env-description:"should log at debug level" env-default:"false"`
	Port                 int32  `env:"PORT" env-description:"server port" env-default:"50051"`
	Payment_Service_Host string `env:"PAYMENT_SERVICE_HOST" env-description:"wallet service host" env-default:"localhost:50052"`
	Driver_Service_Host  string `env:"DRIVER_SERVICE_HOST" env-description:"wallet service host" env-default:"localhost:50052"`
	Project_Id           string `env:"PROJECT_ID" env-description:"firebase project id" env-default:"NO_PROJECT"`
}

func New() (*Config, error) {
	config := Config{
		Production:           true,
		LogDebug:             false,
		Port:                 50051,
		Project_Id:           "NO_PROJECT",
		Payment_Service_Host: "localhost:50052",
	}

	if err := cleanenv.ReadEnv(&config); err != nil {
		return &config, err
	}

	return &config, nil
}
