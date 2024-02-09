//go:build wireinject

package main

import (
	"github.com/dragonfish-tech/go/pkg/logger"
	"github.com/google/wire"
	"github.com/ride-app/marketplace-service/config"
	apihandlers "github.com/ride-app/marketplace-service/internal/api-handlers"
	driverrepository "github.com/ride-app/marketplace-service/internal/repositories/driver"
	locationrepository "github.com/ride-app/marketplace-service/internal/repositories/location"
	statusrepository "github.com/ride-app/marketplace-service/internal/repositories/status"
	walletrepository "github.com/ride-app/marketplace-service/internal/repositories/wallet"
	thirdparty "github.com/ride-app/marketplace-service/third-party"
)

func InitializeService(logger logger.Logger, config *config.Config) (*apihandlers.MarketplaceServiceServer, error) {
	panic(
		wire.Build(
			thirdparty.NewFirebaseApp,
			statusrepository.NewFirebaseStatusRepository,
			wire.Bind(
				new(statusrepository.StatusRepository),
				new(*statusrepository.FirebaseImpl),
			),
			locationrepository.NewFirebaseLocationRepository,
			wire.Bind(
				new(locationrepository.LocationRepository),
				new(*locationrepository.FirebaseImpl),
			),
			driverrepository.New,
			wire.Bind(
				new(driverrepository.DriverRepository),
				new(*driverrepository.Impl),
			),
			walletrepository.New,
			wire.Bind(
				new(walletrepository.WalletRepository),
				new(*walletrepository.Impl),
			),
			apihandlers.New,
		),
	)
}
