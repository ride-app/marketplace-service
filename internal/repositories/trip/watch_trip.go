package triprepository

import (
	"context"
	"encoding/json"

	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WatchTripResult struct {
	Trip *pb.Trip
	Err  error
}

func (r *FirebaseImpl) WatchTrip(
	ctx context.Context,
	log logger.Logger,
	id string,
	watchTripResults chan<- *WatchTripResult,
) {
	iterator := r.firestore.Collection("trips").Doc(id).Snapshots(ctx)
	defer iterator.Stop()

	for {
		snap, err := iterator.Next()
		// DeadlineExceeded will be returned when ctx is cancelled.
		if status.Code(err) == codes.DeadlineExceeded {
			watchTripResults <- &WatchTripResult{
				Trip: nil,
				Err:  err,
			}
			close(watchTripResults)
		}

		if err != nil {
			log.WithError(err).Error("could not get next snapshot")
			watchTripResults <- &WatchTripResult{
				Trip: nil,
				Err:  err,
			}
			close(watchTripResults)
		}

		if !snap.Exists() {
			log.WithError(err).Error("document deleted")
			watchTripResults <- &WatchTripResult{
				Trip: nil,
				Err:  err,
			}
			close(watchTripResults)
		}

		trip := &pb.Trip{}

		jsonBytes, err := json.Marshal(snap.Data())
		if err != nil {
			log.WithError(err).Error("could not marshal trip json into bytes")
			watchTripResults <- &WatchTripResult{
				Trip: nil,
				Err:  err,
			}
			close(watchTripResults)
		}

		err = protojson.Unmarshal(jsonBytes, trip)

		if err != nil {
			log.WithError(err).Error("could not unmarshal trip json bytes into Trip")
			watchTripResults <- &WatchTripResult{
				Trip: nil,
				Err:  err,
			}
			close(watchTripResults)
		}

		trip.CreateTime = timestamppb.New(snap.CreateTime)
		trip.UpdateTime = timestamppb.New(snap.UpdateTime)

		if err != nil {
			log.WithError(err).Error("could not convert snapshot to trip")
			watchTripResults <- &WatchTripResult{
				Trip: nil,
				Err:  err,
			}
			close(watchTripResults)
		}

		watchTripResults <- &WatchTripResult{
			Trip: trip,
			Err:  nil,
		}
	}
}
