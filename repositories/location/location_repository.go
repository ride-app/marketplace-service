//go:generate go run go.uber.org/mock/mockgen -destination ../../testing/mocks/$GOFILE -package mocks . LocationRepository

package locationrepository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/mmcloughlin/geohash"
	pb "github.com/ride-app/marketplace-service/api/gen/ride/marketplace/v1alpha1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LocationRepository interface {
	GetLocation(ctx context.Context, id string) (*pb.Location, error)

	UpdateLocation(ctx context.Context, id string, location *pb.Location) (updateTime *time.Time, err error)
}

type FirebaseImpl struct {
	firestore *firestore.Client
}

func NewFirebaseLocationRepository(firebaseApp *firebase.App) (*FirebaseImpl, error) {
	firestore, err := firebaseApp.Firestore(context.Background())

	if err != nil {
		logrus.WithError(err).Error("Error initializing firestore client")
		return nil, err
	}

	if err != nil {
		logrus.WithError(err).Error("Error initializing auth client")
		return nil, err
	}

	logrus.Info("Firebase location repository initialized")
	return &FirebaseImpl{
		firestore: firestore,
	}, nil
}

func (r *FirebaseImpl) GetLocation(ctx context.Context, id string) (*pb.Location, error) {
	logrus.Info("Checking if driver is active in firestore")
	doc, err := r.firestore.Collection("activeDrivers").Doc(id).Get(ctx)

	if status.Code(err) == codes.NotFound {
		logrus.Info("Driver does not exist in active drivers in firestore")
		return nil, nil
	} else if err != nil {
		logrus.WithError(err).Error("Error checking if driver is active in firestore")
		return nil, err
	}

	if !doc.Exists() {
		logrus.Info("Driver is not active in firestore")
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

func (r *FirebaseImpl) UpdateLocation(ctx context.Context, id string, location *pb.Location) (updateTime *time.Time, err error) {

	logrus.Info("Calculating geohash")
	hash := geohash.Encode(location.Latitude, location.Longitude)

	logrus.Info("Updating driver location in firestore")
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
		logrus.WithError(err).Error("Error updating driver location in firestore")
		return nil, err
	}

	return &res.UpdateTime, nil
}
