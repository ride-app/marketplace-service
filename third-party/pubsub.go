package thirdparty

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/dragonfish/go/v2/pkg/logger"
	"github.com/ride-app/marketplace-service/config"
)

func NewPubSubClient(log logger.Logger, config *config.Config) (*pubsub.Client, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.Project_Id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}
