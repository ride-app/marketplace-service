package service

import (
	dr "github.com/ride-app/marketplace-service/repositories/driver"
	lr "github.com/ride-app/marketplace-service/repositories/location"
	sr "github.com/ride-app/marketplace-service/repositories/status"
	wr "github.com/ride-app/marketplace-service/repositories/wallet"
	"github.com/ride-app/marketplace-service/utils/logger"
)

type MarketplaceServiceServer struct {
	statusRepository   sr.StatusRepository
	locationRepository lr.LocationRepository
	walletrepository   wr.WalletRepository
	driverRepository   dr.DriverRepository
	logger             logger.Logger
}

func New(
	statusRepository sr.StatusRepository,
	locationRepository lr.LocationRepository,
	walletrepository wr.WalletRepository,
	driverRepository dr.DriverRepository,
	logger logger.Logger,
) *MarketplaceServiceServer {
	return &MarketplaceServiceServer{
		statusRepository:   statusRepository,
		locationRepository: locationRepository,
		walletrepository:   walletrepository,
		driverRepository:   driverRepository,
		logger:             logger,
	}
}
