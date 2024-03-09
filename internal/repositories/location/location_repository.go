//go:generate go run go.uber.org/mock/mockgen -destination ./mock/$GOFILE . LocationRepository

package locationrepository

import (
	"context"
	"time"

	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
)

type LocationRepository interface {
	GetLocation(ctx context.Context, log logger.Logger, id string) (*pb.Location, error)

	ListenLocation(
		ctx context.Context,
		log logger.Logger,
		id string,
		locationResponseStream chan<- *LocationStreamResponse,
	)

	UpdateLocation(
		ctx context.Context,
		log logger.Logger,
		id string,
		location *pb.Location,
	) (updateTime *time.Time, err error)
}
