package driverrepository

import (
	"net/http"

	driverApi "buf.build/gen/go/ride/driver/connectrpc/go/ride/driver/v1alpha1/driverv1alpha1connect"
	"github.com/dragonfish/go/v2/pkg/logger"
	"github.com/ride-app/marketplace-service/config"
)

type Impl struct {
	driverApi driverApi.DriverServiceClient
}

func New(log logger.Logger, config *config.Config) (*Impl, error) {
	log.Debug("driver Service Host: ", config.Driver_Service_Host)
	client := driverApi.NewDriverServiceClient(
		http.DefaultClient,
		config.Driver_Service_Host,
	)

	log.Info("driver Repository initialized")
	return &Impl{driverApi: client}, nil
}
