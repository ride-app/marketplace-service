//go:generate go run go.uber.org/mock/mockgen -destination ./mock/$GOFILE . DriverRepository

package driverrepository

import (
	"context"
	"net/http"

	driverApi "buf.build/gen/go/ride/driver/connectrpc/go/ride/driver/v1alpha1/driverv1alpha1connect"
	pb "buf.build/gen/go/ride/driver/protocolbuffers/go/ride/driver/v1alpha1"
	"connectrpc.com/connect"
	"github.com/dragonfish/go/pkg/logger"
	"github.com/ride-app/marketplace-service/config"
)

type DriverRepository interface {
	GetDriver(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Driver, error)

	GetVehicle(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Vehicle, error)
}

type Impl struct {
	driverApi driverApi.DriverServiceClient
}

func New(log logger.Logger, config *config.Config) (*Impl, error) {
	log.Debug("driver Service Host: ", config.Driver_Service_Host)
	client := driverApi.NewDriverServiceClient(
		http.DefaultClient,
		config.Driver_Service_Host,
	)

	log.Info("driver Repository initialized")
	return &Impl{driverApi: client}, nil
}

func (r *Impl) GetDriver(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Driver, error) {
	log.Info("Getting driver from driver service")
	req := connect.NewRequest(&pb.GetDriverRequest{
		Name: "drivers/" + id,
	})
	req.Header().Add("Authorization", authToken)

	res, err := r.driverApi.GetDriver(ctx, req)

	if connect.CodeOf(err) == connect.CodeNotFound {
		log.WithError(err).Error("driver not found")
		return nil, nil
	} else if err != nil {
		log.WithError(err).Error("error getting driver from driver service")
		return nil, err
	}

	return res.Msg.Driver, nil
}

func (r *Impl) GetVehicle(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Vehicle, error) {
	log.Info("getting vehicle from driver service")
	req := connect.NewRequest(&pb.GetVehicleRequest{
		Name: "drivers/" + id + "/vehicle",
	})
	req.Header().Add("Authorization", authToken)

	res, err := r.driverApi.GetVehicle(ctx, req)

	if connect.CodeOf(err) == connect.CodeNotFound {
		log.WithError(err).Error("vehicle not found")
		return nil, nil
	} else if err != nil {
		log.WithError(err).Error("error getting vehicle from driver service")
		return nil, err
	}

	return res.Msg.Vehicle, nil
}
