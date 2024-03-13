package triprepository

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	types "github.com/ride-app/marketplace-service/internal/utils/types"
	"google.golang.org/protobuf/proto"
)

func (r *FirebaseCloudPubSubImpl) WatchTripsCreated(
	ctx context.Context,
	log logger.Logger,
	newTripsResult chan<- *types.StreamResult[*types.Event[*pb.Trip]],
) {
	subscription := r.pubsub.Subscription("trip/created")

	err := subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		trip := &pb.Trip{}

		err := proto.Unmarshal(msg.Data, trip)
		if err != nil {
			newTripsResult <- &types.StreamResult[*types.Event[*pb.Trip]]{
				Result: nil,
				Error:  err,
			}
			msg.Nack()
		}

		newTripsResult <- &types.StreamResult[*types.Event[*pb.Trip]]{
			Result: &types.Event[*pb.Trip]{
				Attributes: msg.Attributes,
				Data:       trip,
				Timestamp:  msg.PublishTime,
			},
			Error: nil,
		}

		msg.Ack()
	})
	if err != nil {
		newTripsResult <- &types.StreamResult[*types.Event[*pb.Trip]]{
			Result: nil,
			Error:  err,
		}
	}
}
