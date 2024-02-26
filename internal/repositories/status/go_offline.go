package statusrepository

import (
	"context"

	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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
