//go:generate go run go.uber.org/mock/mockgen -destination ./mock/$GOFILE . DriverRepository

package driverrepository

import (
	"context"

	pb "buf.build/gen/go/ride/driver/protocolbuffers/go/ride/driver/v1alpha1"
	"github.com/dragonfish/go/v2/pkg/logger"
)

type DriverRepository interface {
	GetDriver(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Driver, error)

	GetVehicle(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Vehicle, error)
}
