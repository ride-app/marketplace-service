package config

type ConfigStruct struct {
	Production          bool   `env:"PRODUCTION" env-description:"dev or prod" env-default:"true"`
	LogDebug            bool   `env:"LOG_DEBUG" env-description:"should log at debug level" env-default:"false"`
	Port                int32  `env:"PORT" env-description:"server port" env-default:"50051"`
	Wallet_Service_Host string `env:"WALLET_SERVICE_HOST" env-description:"wallet service host" env-default:"localhost:50052"`
	Driver_Service_Host string `env:"DRIVER_SERVICE_HOST" env-description:"wallet service host" env-default:"localhost:50052"`
	Project_Id          string `env:"PROJECT_ID" env-description:"firebase project id" env-default:"NO_PROJECT"`
}

var Env ConfigStruct
