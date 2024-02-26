package triprepository

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/protobuf/encoding/protojson"
)

func (r *FirebaseCloudPubSubImpl) UpdateTrip(ctx context.Context, log logger.Logger, trip *pb.Trip) (updateTime *time.Time, err error) {
	trip.CreateTime = nil
	trip.UpdateTime = nil

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

	return &writeResult.UpdateTime, nil
}
