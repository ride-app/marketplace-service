package locationrepository

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/dragonfish/go/v2/pkg/logger"
)

type FirebaseImpl struct {
	firestore *firestore.Client
}

func NewFirebaseLocationRepository(
	log logger.Logger,
	firebaseApp *firebase.App,
) (*FirebaseImpl, error) {
	firestore, err := firebaseApp.Firestore(context.Background())
	if err != nil {
		log.WithError(err).Error("error initializing firestore client")
		return nil, err
	}

	if err != nil {
		log.WithError(err).Error("error initializing auth client")
		return nil, err
	}

	log.Info("firebase location repository initialized")
	return &FirebaseImpl{
		firestore: firestore,
	}, nil
}
