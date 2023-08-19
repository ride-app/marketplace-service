//go:generate go run go.uber.org/mock/mockgen -destination ../../testing/mocks/$GOFILE -package mocks . DriverRepository

package driverrepository

import (
	"context"
	"net/http"

	driverApi "buf.build/gen/go/ride/driver/connectrpc/go/ride/driver/v1alpha1/driverv1alpha1connect"
	pb "buf.build/gen/go/ride/driver/protocolbuffers/go/ride/driver/v1alpha1"
	"connectrpc.com/connect"
	"github.com/ride-app/marketplace-service/config"
	"github.com/ride-app/marketplace-service/utils/logger"
)

type DriverRepository interface {
	GetDriver(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Driver, error)

	GetVehicle(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Vehicle, error)
}

type Impl struct {
	driverApi driverApi.DriverServiceClient
}

func New(log logger.Logger, config *config.Config) (*Impl, error) {
	log.Debug("Driver Service Host: ", config.Driver_Service_Host)
	client := driverApi.NewDriverServiceClient(
		http.DefaultClient,
		config.Driver_Service_Host,
	)

	log.Info("Driver Repository initialized")
	return &Impl{driverApi: client}, nil
}

func (r *Impl) GetDriver(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Driver, error) {
	log.Info("Getting driver from driver service")
	req := connect.NewRequest(&pb.GetDriverRequest{
		Name: "driver/" + id,
	})
	req.Header().Add("Authorization", authToken)

	res, err := r.driverApi.GetDriver(ctx, req)

	if connect.CodeOf(err) == connect.CodeNotFound {
		log.WithError(err).Error("Driver not found")
		return nil, nil
	} else if err != nil {
		log.WithError(err).Error("Error getting driver from driver service")
		return nil, err
	}

	return res.Msg.Driver, nil
}

func (r *Impl) GetVehicle(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Vehicle, error) {
	log.Info("Getting vehicle from driver service")
	req := connect.NewRequest(&pb.GetVehicleRequest{
		Name: "driver/" + id + "/vehicle",
	})
	req.Header().Add("Authorization", authToken)

	res, err := r.driverApi.GetVehicle(ctx, req)

	if connect.CodeOf(err) == connect.CodeNotFound {
		log.WithError(err).Error("Vehicle not found")
		return nil, nil
	} else if err != nil {
		log.WithError(err).Error("Error getting vehicle from driver service")
		return nil, err
	}

	return res.Msg.Vehicle, nil
}
