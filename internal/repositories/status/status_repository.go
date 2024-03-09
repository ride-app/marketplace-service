//go:generate go run go.uber.org/mock/mockgen -destination ./mock/$GOFILE . StatusRepository

package statusrepository

import (
	"context"

	driverv1alpha1 "buf.build/gen/go/ride/driver/protocolbuffers/go/ride/driver/v1alpha1"
	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
)

type StatusRepository interface {
	GetStatus(ctx context.Context, log logger.Logger, id string) (*pb.Status, error)

	ListenStatus(
		ctx context.Context,
		log logger.Logger,
		id string,
		statusResponseStream chan<- *StatusStreamResponse,
	)

	GoOnline(
		ctx context.Context,
		log logger.Logger,
		id string,
		vehicleType *driverv1alpha1.Vehicle,
	) (*pb.Status, error)

	GoOffline(ctx context.Context, log logger.Logger, id string) (*pb.Status, error)
}
