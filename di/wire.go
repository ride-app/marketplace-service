//go:build wireinject

package di

import (
	"github.com/google/wire"
	driverrepository "github.com/ride-app/marketplace-service/repositories/driver"
	locationrepository "github.com/ride-app/marketplace-service/repositories/location"
	statusrepository "github.com/ride-app/marketplace-service/repositories/status"
	walletrepository "github.com/ride-app/marketplace-service/repositories/wallet"
	"github.com/ride-app/marketplace-service/service"
	thirdparty "github.com/ride-app/marketplace-service/third-party"
)

func InitializeService() (*service.MarketplaceServiceServer, error) {
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
			service.New,
		),
	)
}
