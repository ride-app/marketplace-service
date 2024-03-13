package apihandlers

import (
	"github.com/dragonfish/go/v2/pkg/logger"
	dr "github.com/ride-app/marketplace-service/internal/repositories/driver"
	lr "github.com/ride-app/marketplace-service/internal/repositories/location"
	sr "github.com/ride-app/marketplace-service/internal/repositories/status"
	tr "github.com/ride-app/marketplace-service/internal/repositories/trip"
	wr "github.com/ride-app/marketplace-service/internal/repositories/wallet"
)

type MarketplaceServiceServer struct {
	statusRepository   sr.StatusRepository
	locationRepository lr.LocationRepository
	walletrepository   wr.WalletRepository
	driverRepository   dr.DriverRepository
	tripRepository     tr.TripRepository
	logger             logger.Logger
}

func New(
	statusRepository sr.StatusRepository,
	locationRepository lr.LocationRepository,
	walletrepository wr.WalletRepository,
	driverRepository dr.DriverRepository,
	tripRepository tr.TripRepository,
	logger logger.Logger,
) *MarketplaceServiceServer {
	return &MarketplaceServiceServer{
		statusRepository:   statusRepository,
		locationRepository: locationRepository,
		walletrepository:   walletrepository,
		driverRepository:   driverRepository,
		tripRepository:     tripRepository,
		logger:             logger,
	}
}
