package triprepository

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func (r *FirebaseCloudPubSubImpl) CreateTrip(ctx context.Context, log logger.Logger, trip *pb.Trip) (createTime *time.Time, err error) {
	trip.Status = pb.Trip_STATUS_PENDING
	trip.Driver = nil
	trip.Route.WalkToPickup = nil
	trip.Route.WalkToDestination = nil
	trip.CreateTime = nil
	trip.UpdateTime = nil
	trip.StartTime = nil
	trip.EndTime = nil

	jsonData, err := protojson.Marshal(trip)

	if err != nil {
		log.WithError(err).Error("could not marshal trip into json")
		return nil, err
	}

	doc := make(map[string]interface{})

	err = json.Unmarshal(jsonData, &doc)

	if err != nil {
		log.WithError(err).Error("could not unmarshal trip json into map")
		return nil, err
	}

	id := strings.Split(trip.Name, "/")[1]
	writeResult, err := r.firestore.Collection("trips").Doc(id).Set(ctx, doc)

	if err != nil {
		log.WithError(err).Error("could not write trip data to firestore")
		return nil, err
	}

	go func() {
		data, err := proto.Marshal(trip)

		if err != nil {
			log.WithError(err).Warn("could not marshal trip into []byte")
		}

		msg := &pubsub.Message{}
		msg.Data = data

		r.pubsub.Topic("trips/created").Publish(ctx, msg)
	}()

	return &writeResult.UpdateTime, nil
}
