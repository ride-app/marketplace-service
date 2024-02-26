package locationrepository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/dragonfish/go/v2/pkg/logger"
	"github.com/mmcloughlin/geohash"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
)

func (r *FirebaseImpl) UpdateLocation(ctx context.Context, log logger.Logger, id string, location *pb.Location) (updateTime *time.Time, err error) {
	log.Info("calculating geohash")
	hash := geohash.Encode(location.Latitude, location.Longitude)

	log.Info("updating driver location in firestore")
	res, err := r.firestore.Collection("activeDrivers").Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "location.latitude",
			Value: location.Latitude,
		},
		{
			Path:  "location.longitude",
			Value: location.Longitude,
		},
		{
			Path:  "geohash",
			Value: hash,
		},
	})

	if err != nil {
		log.WithError(err).Error("error updating driver location in firestore")
		return nil, err
	}

	return &res.UpdateTime, nil
}
