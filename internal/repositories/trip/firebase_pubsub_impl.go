package triprepository

import (
	"context"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/dragonfish/go/v2/pkg/logger"
)

type FirebaseCloudPubSubImpl struct {
	firestore *firestore.Client
	auth      *auth.Client
	pubsub    *pubsub.Client
}

func NewFirebaseTripRepository(
	log logger.Logger,
	firebaseApp *firebase.App,
	auth *auth.Client,
	pubsub *pubsub.Client,
) (*FirebaseCloudPubSubImpl, error) {
	firestore, err := firebaseApp.Firestore(context.Background())
	if err != nil {
		log.WithError(err).Error("error initializing firestore client")
		return nil, err
	}

	if err != nil {
		log.WithError(err).Error("error initializing auth client")
		return nil, err
	}

	log.Info("firebase trip repository initialized")
	return &FirebaseCloudPubSubImpl{
		firestore: firestore,
		auth:      auth,
	}, nil
}
