//go:generate go run go.uber.org/mock/mockgen -destination ./mock/$GOFILE . StatusRepository

package statusrepository

import (
	"context"
	"fmt"
	"strings"

	driverv1alpha1 "buf.build/gen/go/ride/driver/protocolbuffers/go/ride/driver/v1alpha1"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/dragonfish/go/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type StatusRepository interface {
	GetStatus(ctx context.Context, log logger.Logger, id string) (*pb.Status, error)

	ListenStatus(ctx context.Context, log logger.Logger, id string, statusResponseStream chan<- *StatusStreamResponse)

	GoOnline(ctx context.Context, log logger.Logger, id string, vehicleType *driverv1alpha1.Vehicle) (*pb.Status, error)

	GoOffline(ctx context.Context, log logger.Logger, id string) (*pb.Status, error)
}

type FirebaseImpl struct {
	firestore *firestore.Client

	capacities map[driverv1alpha1.Vehicle_Type]int
}

func NewFirebaseStatusRepository(log logger.Logger, firebaseApp *firebase.App) (*FirebaseImpl, error) {
	firestore, err := firebaseApp.Firestore(context.Background())

	if err != nil {
		log.WithError(err).Error("error initializing firestore client")
		return nil, err
	}

	if err != nil {
		log.WithError(err).Error("error initializing auth client")
		return nil, err
	}

	log.Info("firebase status repository initialized")
	return &FirebaseImpl{
		firestore: firestore,
		capacities: map[driverv1alpha1.Vehicle_Type]int{
			driverv1alpha1.Vehicle_TYPE_AUTORICKSHAW: 4,
			driverv1alpha1.Vehicle_TYPE_ERICKSHAW:    4,
			driverv1alpha1.Vehicle_TYPE_MOTORCYCLE:   1,
		},
	}, nil
}

func (r *FirebaseImpl) GetStatus(ctx context.Context, log logger.Logger, id string) (*pb.Status, error) {
	log.Info("getting status from firestore")
	doc, err := r.firestore.Collection("activeDrivers").Doc(id).Get(ctx)

	if status.Code(err) == codes.NotFound {
		log.Info("driver does not exist in firestore")
		return nil, nil
	} else if err != nil {
		log.WithError(err).Error("error getting status from firestore")
		return nil, err
	}

	if !doc.Exists() {
		log.Info("driver does not exist in firestore")
		return nil, nil
	}

	status := pb.Status{
		Name:       "drivers/" + id + "/status",
		Online:     doc.Exists(),
		UpdateTime: timestamppb.New(doc.UpdateTime),
	}

	return &status, nil
}

type StatusStreamResponse struct {
	Status *pb.Status
	Error  error
}

func (r *FirebaseImpl) ListenStatus(ctx context.Context, log logger.Logger, id string, statusResponseStream chan<- *StatusStreamResponse) {
	log.Info("listening for status updates from firestore")
	snapshots := r.firestore.Collection("activeDrivers").Doc(id).Snapshots(ctx)
	defer snapshots.Stop()

	for {
		log.Info("waiting for status update from firestore")
		snap, err := snapshots.Next()
		log.Info("got status update from firestore")

		if status.Code(err) == codes.DeadlineExceeded {
			log.Info("firestore deadline exceeded")
			statusResponseStream <- &StatusStreamResponse{
				Status: nil,
				Error:  nil,
			}
		} else if status.Code(err) == codes.NotFound {
			log.Info("driver does not exist in firestore")
			statusResponseStream <- &StatusStreamResponse{
				Status: nil,
				Error:  nil,
			}
		} else if err != nil {
			log.WithError(err).Error("error getting status from firestore")
			statusResponseStream <- &StatusStreamResponse{
				Status: nil,
				Error:  err,
			}
		}

		if !snap.Exists() {
			log.Info("driver does not exist in firestore")
			statusResponseStream <- &StatusStreamResponse{
				Status: nil,
				Error:  fmt.Errorf("document %s does not exist", snap.Ref.ID),
			}
		}

		statusResponseStream <- &StatusStreamResponse{
			Status: &pb.Status{
				Name:       "drivers/" + id + "/status",
				Online:     snap.Exists(),
				UpdateTime: timestamppb.New(snap.UpdateTime),
			},
			Error: nil,
		}
	}
}

func (r *FirebaseImpl) GoOnline(ctx context.Context, log logger.Logger, id string, vehicle *driverv1alpha1.Vehicle) (*pb.Status, error) {
	log.Info("ppdating active driver in firestore")

	ref := r.firestore.Collection("activeDrivers").Doc(id)
	err := r.firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(ref)

		if err != nil && !(status.Code(err) == codes.NotFound) {
			log.WithError(err).Error("error getting active driver from firestore")
			return err
		}

		if doc.Exists() {
			return nil
		}

		err = tx.Set(ref, map[string]interface{}{
			"vehicleId":    strings.Split(vehicle.Name, "/")[1],
			"licensePlate": vehicle.LicensePlate,
			"vehicleType":  strings.ToLower(vehicle.Type.String()),
			"capacity":     r.capacities[vehicle.Type],
		})

		if err != nil {
			log.WithError(err).Error("error setting active driver in firestore")
			return err
		}

		return nil
	})

	if err != nil {
		log.WithError(err).Error("error updating active driver in firestore")
		return nil, err
	}

	return &pb.Status{
		Name:       "drivers/" + id + "/status",
		Online:     true,
		UpdateTime: timestamppb.Now(),
	}, nil
}

func (r *FirebaseImpl) GoOffline(ctx context.Context, log logger.Logger, id string) (*pb.Status, error) {
	log.Info("deleting active driver from firestore")
	_, err := r.firestore.Collection("activeDrivers").Doc(id).Delete(ctx)

	if status.Code(err) == codes.NotFound {
		log.Info("driver does not exist in active drivers in firestore")
	} else if err != nil {
		log.WithError(err).Error("error deleting active driver from firestore")
		return nil, err
	}

	return &pb.Status{
		Name:       "drivers/" + id + "/status",
		Online:     false,
		UpdateTime: timestamppb.Now(),
	}, nil
}
