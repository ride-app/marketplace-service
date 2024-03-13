package statusrepository

import (
	"context"
	"fmt"

	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type StatusStreamResponse struct {
	Status *pb.Status
	Error  error
}

func (r *FirebaseImpl) ListenStatus(
	ctx context.Context,
	log logger.Logger,
	id string,
	statusResponseStream chan<- *StatusStreamResponse,
) {
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
