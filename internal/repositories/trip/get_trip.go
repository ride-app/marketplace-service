package triprepository

import (
	"context"
	"encoding/json"

	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (r *FirebaseCloudPubSubImpl) GetTrip(ctx context.Context, log logger.Logger, id string) (*pb.Trip, error) {
	log.Info("querying trip from firestore")
	snap, err := r.firestore.Collection("trips").Doc(id).Get(ctx)

	if err != nil {
		log.WithError(err).Error("error querying trip from firestore")
		return nil, err
	}

	trip := &pb.Trip{}

	jsonBytes, err := json.Marshal(snap.Data())

	if err != nil {
		log.WithError(err).Error("could not marshal trip json into bytes")
		return nil, err
	}

	err = protojson.Unmarshal(jsonBytes, trip)

	if err != nil {
		log.WithError(err).Error("could not unmarshal trip json bytes into Trip")
		return nil, err
	}

	trip.CreateTime = timestamppb.New(snap.CreateTime)
	trip.UpdateTime = timestamppb.New(snap.UpdateTime)

	// trip, err := r.docToTrip(ctx, log, snap)

	// if err != nil {
	// 	log.WithError(err).Error("could not convert firebase document to trip")
	// }

	return trip, nil
}
