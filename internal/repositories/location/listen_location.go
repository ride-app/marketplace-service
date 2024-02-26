package locationrepository

import (
	"context"

	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
