//go:generate go run go.uber.org/mock/mockgen -destination ./mock/$GOFILE . DriverRepository

package driverrepository

import (
	"context"

	pb "buf.build/gen/go/ride/driver/protocolbuffers/go/ride/driver/v1alpha1"
	"connectrpc.com/connect"
	"github.com/dragonfish/go/v2/pkg/logger"
)

func (r *Impl) GetDriver(
	ctx context.Context,
	log logger.Logger,
	id string,
	authToken string,
) (*pb.Driver, error) {
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
