package statusrepository

import (
	"context"

	driverv1alpha1 "buf.build/gen/go/ride/driver/protocolbuffers/go/ride/driver/v1alpha1"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/dragonfish/go/v2/pkg/logger"
)

type FirebaseImpl struct {
	firestore *firestore.Client

	capacities map[driverv1alpha1.Vehicle_Type]int
}

func NewFirebaseStatusRepository(
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

	log.Info("firebase status repository initialized")
	return &FirebaseImpl{
		firestore: firestore,
		capacities: map[driverv1alpha1.Vehicle_Type]int{
			driverv1alpha1.Vehicle_TYPE_AUTORICKSHAW: 4,
			driverv1alpha1.Vehicle_TYPE_ERICKSHAW:    4,
			driverv1alpha1.Vehicle_TYPE_MOTORCYCLE:   1,
		},
	}, nil
}
