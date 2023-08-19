package service

import (
	dr "github.com/ride-app/marketplace-service/repositories/driver"
	lr "github.com/ride-app/marketplace-service/repositories/location"
	sr "github.com/ride-app/marketplace-service/repositories/status"
	wr "github.com/ride-app/marketplace-service/repositories/wallet"
)

type MarketplaceServiceServer struct {
	statusRepository   sr.StatusRepository
	locationRepository lr.LocationRepository
	walletrepository   wr.WalletRepository
	driverRepository   dr.DriverRepository
}

func New(
	statusRepository sr.StatusRepository,
	locationRepository lr.LocationRepository,
	walletrepository wr.WalletRepository,
	driverRepository dr.DriverRepository,
) *MarketplaceServiceServer {
	return &MarketplaceServiceServer{
		statusRepository:   statusRepository,
		locationRepository: locationRepository,
		walletrepository:   walletrepository,
		driverRepository:   driverRepository,
	}
}
