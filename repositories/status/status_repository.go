//go:generate go run go.uber.org/mock/mockgen -destination ../../mocks/$GOFILE -package mocks . StatusRepository

package statusrepository

import (
	"context"
	"strings"

	driverv1alpha1 "buf.build/gen/go/ride/driver/protocolbuffers/go/ride/driver/v1alpha1"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	pb "github.com/ride-app/marketplace-service/api/gen/ride/marketplace/v1alpha1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type StatusRepository interface {
	GetStatus(ctx context.Context, id string) (*pb.Status, error)

	GoOnline(ctx context.Context, id string, vehicleType *driverv1alpha1.Vehicle) (*pb.Status, error)

	GoOffline(ctx context.Context, id string) (*pb.Status, error)
}

type FirebaseImpl struct {
	firestore *firestore.Client
}

func NewFirebaseStatusRepository(firebaseApp *firebase.App) (*FirebaseImpl, error) {
	firestore, err := firebaseApp.Firestore(context.Background())

	if err != nil {
		logrus.WithError(err).Error("Error initializing firestore client")
		return nil, err
	}

	if err != nil {
		logrus.WithError(err).Error("Error initializing auth client")
		return nil, err
	}

	logrus.Info("Firebase status repository initialized")
	return &FirebaseImpl{
		firestore: firestore,
	}, nil
}

func (r *FirebaseImpl) GetStatus(ctx context.Context, id string) (*pb.Status, error) {
	logrus.Info("Getting status from firestore")
	doc, err := r.firestore.Collection("activeDrivers").Doc(id).Get(ctx)

	if status.Code(err) == codes.NotFound {
		logrus.Info("Driver does not exist in firestore")
		return nil, nil
	} else if err != nil {
		logrus.WithError(err).Error("Error getting status from firestore")
		return nil, err
	}

	if !doc.Exists() {
		logrus.Info("Driver does not exist in firestore")
		return nil, nil
	}

	status := pb.Status{
		Name:       "drivers/" + id + "/status",
		Online:     doc.Exists(),
		UpdateTime: timestamppb.New(doc.UpdateTime),
	}

	return &status, nil
}

func (r *FirebaseImpl) GoOnline(ctx context.Context, id string, vehicle *driverv1alpha1.Vehicle) (*pb.Status, error) {
	logrus.Info("Updating active driver in firestore")

	ref := r.firestore.Collection("activeDrivers").Doc(id)
	err := r.firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(ref)

		if err != nil && !(status.Code(err) == codes.NotFound) {
			logrus.WithError(err).Error("Error getting active driver from firestore")
			return err
		}

		if doc.Exists() {
			return nil
		}

		err = tx.Set(ref, map[string]interface{}{
			"vehicleId":    strings.Split(vehicle.Name, "/")[1],
			"licensePlate": vehicle.LicensePlate,
			"vehicleType":  strings.ToLower(vehicle.Type.String()),
			"capacity":     4,
		})

		if err != nil {
			logrus.WithError(err).Error("Error setting active driver in firestore")
			return err
		}

		return nil
	})

	if err != nil {
		logrus.WithError(err).Error("Error updating active driver in firestore")
		return nil, err
	}

	return &pb.Status{
		Name:       "drivers/" + id + "/status",
		Online:     true,
		UpdateTime: timestamppb.Now(),
	}, nil
}

func (r *FirebaseImpl) GoOffline(ctx context.Context, id string) (*pb.Status, error) {
	logrus.Info("Deleting active driver from firestore")
	_, err := r.firestore.Collection("activeDrivers").Doc(id).Delete(ctx)

	if status.Code(err) == codes.NotFound {
		logrus.Info("Driver does not exist in active drivers in firestore")
	} else if err != nil {
		logrus.WithError(err).Error("Error deleting active driver from firestore")
		return nil, err
	}

	return &pb.Status{
		Name:       "drivers/" + id + "/status",
		Online:     false,
		UpdateTime: timestamppb.Now(),
	}, nil
}
