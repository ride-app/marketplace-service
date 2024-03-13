package locationrepository

import (
	"context"

	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *FirebaseImpl) GetLocation(
	ctx context.Context,
	log logger.Logger,
	id string,
) (*pb.Location, error) {
	log.Info("checking if driver is active in firestore")
	doc, err := r.firestore.Collection("activeDrivers").Doc(id).Get(ctx)

	if status.Code(err) == codes.NotFound {
		log.Info("driver does not exist in active drivers in firestore")
		return nil, nil
	} else if err != nil {
		log.WithError(err).Error("error checking if driver is active in firestore")
		return nil, err
	}

	if !doc.Exists() {
		log.Info("driver is not active in firestore")
		return nil, nil
	}

	data := doc.Data()

	location := data["location"].(map[string]interface{})
	latitude := location["latitude"].(float64)
	longitude := location["longitude"].(float64)

	return &pb.Location{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}
