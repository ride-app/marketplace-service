//go:generate go run go.uber.org/mock/mockgen -destination ../../testing/mocks/$GOFILE -package mocks . LocationRepository

package locationrepository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/mmcloughlin/geohash"
	pb "github.com/ride-app/marketplace-service/api/gen/ride/marketplace/v1alpha1"
	"github.com/ride-app/marketplace-service/utils/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LocationRepository interface {
	GetLocation(ctx context.Context, log logger.Logger, id string) (*pb.Location, error)

	ListenLocation(ctx context.Context, log logger.Logger, id string, locationResponseStream chan<- *LocationStreamResponse)

	UpdateLocation(ctx context.Context, log logger.Logger, id string, location *pb.Location) (updateTime *time.Time, err error)
}

type FirebaseImpl struct {
	firestore *firestore.Client
}

func NewFirebaseLocationRepository(log logger.Logger, firebaseApp *firebase.App) (*FirebaseImpl, error) {
	firestore, err := firebaseApp.Firestore(context.Background())

	if err != nil {
		log.WithError(err).Error("error initializing firestore client")
		return nil, err
	}

	if err != nil {
		log.WithError(err).Error("error initializing auth client")
		return nil, err
	}

	log.Info("firebase location repository initialized")
	return &FirebaseImpl{
		firestore: firestore,
	}, nil
}

func (r *FirebaseImpl) GetLocation(ctx context.Context, log logger.Logger, id string) (*pb.Location, error) {
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

type LocationStreamResponse struct {
	Location *pb.Location
	Error    error
}

func (r *FirebaseImpl) ListenLocation(ctx context.Context, log logger.Logger, id string, locationResponseStream chan<- *LocationStreamResponse) {
	log.Info("listening for driver location changes in firestore")
	snapshots := r.firestore.Collection("activeDrivers").Doc(id).Snapshots(ctx)

	defer snapshots.Stop()

	for {
		snapshot, err := snapshots.Next()

		if status.Code(err) == codes.DeadlineExceeded {
			log.Info("firestore deadline exceeded")
			locationResponseStream <- &LocationStreamResponse{
				Location: nil,
				Error:    nil,
			}
		} else if status.Code(err) == codes.NotFound {
			log.Info("driver does not exist in firestore")
			locationResponseStream <- &LocationStreamResponse{
				Location: nil,
				Error:    nil,
			}
		} else if err != nil {
			log.WithError(err).Error("error listening for driver location changes from firestore")
			locationResponseStream <- &LocationStreamResponse{
				Location: nil,
				Error:    err,
			}
		}

		if !snapshot.Exists() {
			log.Info("driver location not found in firestore")
			locationResponseStream <- &LocationStreamResponse{
				Location: nil,
				Error:    nil,
			}
			return
		}

		data := snapshot.Data()

		location := data["location"].(map[string]interface{})
		latitude := location["latitude"].(float64)
		longitude := location["longitude"].(float64)

		locationResponseStream <- &LocationStreamResponse{
			Location: &pb.Location{
				Latitude:  latitude,
				Longitude: longitude,
			},
			Error: nil,
		}
	}
}

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
