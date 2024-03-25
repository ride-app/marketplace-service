package triprepository

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/dragonfish/go/v2/pkg/logger"
)

type FirebaseImpl struct {
	firestore *firestore.Client
	auth      *auth.Client
}

func NewFirebaseTripRepository(
	log logger.Logger,
	firebaseApp *firebase.App,
) (*FirebaseImpl, error) {
	firestore, err := firebaseApp.Firestore(context.Background())
	if err != nil {
		log.WithError(err).Error("error initializing firestore client")
		return nil, err
	}

	auth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.WithError(err).Error("error initializing auth client")
		return nil, err
	}

	log.Info("firebase trip repository initialized")
	return &FirebaseImpl{
		firestore: firestore,
		auth:      auth,
	}, nil
}
