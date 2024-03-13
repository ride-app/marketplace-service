package statusrepository

import (
	"context"
	"strings"

	driverv1alpha1 "buf.build/gen/go/ride/driver/protocolbuffers/go/ride/driver/v1alpha1"
	"cloud.google.com/go/firestore"
	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (r *FirebaseImpl) GoOnline(
	ctx context.Context,
	log logger.Logger,
	id string,
	vehicle *driverv1alpha1.Vehicle,
) (*pb.Status, error) {
	log.Info("ppdating active driver in firestore")

	ref := r.firestore.Collection("activeDrivers").Doc(id)
	err := r.firestore.RunTransaction(
		ctx,
		func(ctx context.Context, tx *firestore.Transaction) error {
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
		},
	)
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
