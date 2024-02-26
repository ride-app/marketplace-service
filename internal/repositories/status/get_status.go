package statusrepository

import (
	"context"

	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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
