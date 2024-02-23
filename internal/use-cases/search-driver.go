package usecases

import (
	"context"

	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
)

func SearchDriver(ctx context.Context, log logger.Logger, initialSearchRadius float32, maxSearchRadius float32, req *pb.CreateTripRequest) {
	log.Info("initial search radius in meters", initialSearchRadius)
	log.Info("max search radius in meters", maxSearchRadius)

}
