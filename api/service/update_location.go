package service

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	pb "github.com/ride-app/marketplace-service/api/gen/ride/marketplace/v1alpha1"
)

func (service *MarketplaceServiceServer) UpdateLocation(ctx context.Context,
	req *connect.Request[pb.UpdateLocationRequest]) (*connect.Response[pb.UpdateLocationResponse], error) {
	log := service.logger.WithFields(map[string]string{
		"method": "UpdateLocation",
	})

	if err := req.Msg.Validate(); err != nil {
		log.WithError(err).Info("invalid request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.Parent, "/")[1]

	log.Debug("uid: ", uid)
	log.Debug("request header uid: ", req.Header().Get("uid"))

	if uid != req.Header().Get("uid") {
		log.Info("permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	status, err := service.statusRepository.GetStatus(ctx, log, uid)

	if err != nil {
		log.WithError(err).Error("failed to get status")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !status.Online {
		log.Info("driver is offline")
		return nil, connect.NewError(connect.CodeFailedPrecondition, err)
	}

	_, err = service.locationRepository.UpdateLocation(ctx, log, uid, req.Msg.Location)

	if err != nil {
		log.WithError(err).Error("failed to update location")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("location updated")
	return connect.NewResponse(&pb.UpdateLocationResponse{}), nil
}
