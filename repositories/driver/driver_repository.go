//go:generate go run go.uber.org/mock/mockgen -destination ../../mocks/$GOFILE -package mocks . DriverRepository

package walletrepository

import (
	"context"
	"net/http"

	driverApi "buf.build/gen/go/ride/driver/bufbuild/connect-go/ride/driver/v1alpha1/driverv1alpha1connect"
	pb "buf.build/gen/go/ride/driver/protocolbuffers/go/ride/driver/v1alpha1"
	"github.com/bufbuild/connect-go"
	"github.com/ride-app/marketplace-service/config"
	"github.com/sirupsen/logrus"
)

type DriverRepository interface {
	GetDriver(ctx context.Context, id string, authToken string) (*pb.Driver, error)

	GetVehicle(ctx context.Context, id string, authToken string) (*pb.Vehicle, error)
}

type Impl struct {
	driverApi driverApi.DriverServiceClient
}

func New() (*Impl, error) {
	logrus.Debug("Driver Service Host: ", config.Env.Driver_Service_Host)
	client := driverApi.NewDriverServiceClient(
		http.DefaultClient,
		config.Env.Driver_Service_Host,
	)

	logrus.Info("Driver Repository initialized")
	return &Impl{driverApi: client}, nil
}

func (r *Impl) GetDriver(ctx context.Context, id string, authToken string) (*pb.Driver, error) {
	logrus.Info("Getting driver from driver service")
	req := connect.NewRequest(&pb.GetDriverRequest{
		Name: "driver/" + id,
	})
	req.Header().Add("Authorization", authToken)

	res, err := r.driverApi.GetDriver(ctx, req)

	if connect.CodeOf(err) == connect.CodeNotFound {
		logrus.WithError(err).Error("Driver not found")
		return nil, nil
	} else if err != nil {
		logrus.WithError(err).Error("Error getting driver from driver service")
		return nil, err
	}

	return res.Msg.Driver, nil
}

func (r *Impl) GetVehicle(ctx context.Context, id string, authToken string) (*pb.Vehicle, error) {
	logrus.Info("Getting vehicle from driver service")
	req := connect.NewRequest(&pb.GetVehicleRequest{
		Name: "driver/" + id + "/vehicle",
	})
	req.Header().Add("Authorization", authToken)

	res, err := r.driverApi.GetVehicle(ctx, req)

	if connect.CodeOf(err) == connect.CodeNotFound {
		logrus.WithError(err).Error("Vehicle not found")
		return nil, nil
	} else if err != nil {
		logrus.WithError(err).Error("Error getting vehicle from driver service")
		return nil, err
	}

	return res.Msg.Vehicle, nil
}
