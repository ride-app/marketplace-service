package triprepository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/dragonfish/go/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
)

type TripRepository interface {
	GetTrip(ctx context.Context, log logger.Logger, id string) (*pb.Trip, error)

	CreateTrip(ctx context.Context, log logger.Logger, trip *pb.Trip) (createTime *time.Time, err error)

	ListenTrip(ctx context.Context, log logger.Logger, id string, locationResponseStream chan<- *pb.Trip)

	UpdateTrip(ctx context.Context, log logger.Logger, id string, trip *pb.Trip) (updateTime *time.Time, err error)
}

type FirebaseImpl struct {
	firestore *firestore.Client
}

func NewFirebaseTripRepository(log logger.Logger, firebaseApp *firebase.App) (*FirebaseImpl, error) {
	firestore, err := firebaseApp.Firestore(context.Background())

	if err != nil {
		log.WithError(err).Error("error initializing firestore client")
		return nil, err
	}

	if err != nil {
		log.WithError(err).Error("error initializing auth client")
		return nil, err
	}

	log.Info("firebase trip repository initialized")
	return &FirebaseImpl{
		firestore: firestore,
	}, nil
}

func (r *FirebaseImpl) GetTrip(ctx context.Context, log logger.Logger, id string) (*pb.Trip, error) {
	trip := &pb.Trip{}

	log.Info("querying trip from firestore")
	doc, err := r.firestore.Collection("trips").Doc(id).Get(ctx)

	if err != nil {
		log.WithError(err).Error("error querying trip from firestore")
		return nil, err
	}
	log.Debug("Document data: ", doc.Data())

	if !doc.Exists() {
		log.Warn("trip not found")
		return nil, nil
	}

	log.Info("checking if trip has a driver assigned")
	driverAssigned := doc.Data()["driver.uid"] != nil

	if driverAssigned {
		log.Info("trip has a driver assigned")
	}

	return trip, nil
}

func (r *FirebaseImpl) CreateTrip(ctx context.Context, log logger.Logger, trip *pb.Trip) (string, error) {
}

func (r *FirebaseImpl) UpdateTrip(ctx context.Context, trip *pb.Trip) error {
	_, err := r.firestore.Collection("trips").Doc(trip.Name).Set(ctx, trip)
	return err
}

func (r *FirebaseImpl) DeleteTrip(ctx context.Context, tripID string) error {
	_, err := r.firestore.Collection("trips").Doc(tripID).Delete(ctx)
	return err
}
