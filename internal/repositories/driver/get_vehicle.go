package driverrepository

import (
	"context"

	pb "buf.build/gen/go/ride/driver/protocolbuffers/go/ride/driver/v1alpha1"
	"connectrpc.com/connect"
	"github.com/dragonfish/go/v2/pkg/logger"
)

func (r *Impl) GetVehicle(
	ctx context.Context,
	log logger.Logger,
	id string,
	authToken string,
) (*pb.Vehicle, error) {
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
