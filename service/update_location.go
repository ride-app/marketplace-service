package service

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/marketplace-service/api/gen/ride/marketplace/v1alpha1"
	"github.com/sirupsen/logrus"
)

func (service *MarketplaceServiceServer) UpdateLocation(ctx context.Context,
	req *connect.Request[pb.UpdateLocationRequest]) (*connect.Response[pb.UpdateLocationResponse], error) {
	if err := req.Msg.Validate(); err != nil {
		logrus.Info("Invalid request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.Parent, "/")[1]

	logrus.Debug("uid: ", uid)
	logrus.Debug("Request header uid: ", req.Header().Get("uid"))

	if uid != req.Header().Get("uid") {
		logrus.Info("Permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	status, err := service.statusRepository.GetStatus(ctx, uid)

	if err != nil {
		logrus.WithError(err).Error("Failed to get status")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !status.Online {
		logrus.Info("Driver is offline")
		return nil, connect.NewError(connect.CodeFailedPrecondition, err)
	}

	_, err = service.locationRepository.UpdateLocation(ctx, uid, req.Msg.Location)

	if err != nil {
		logrus.WithError(err).Error("Failed to update location")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	logrus.Info("Location updated")
	return connect.NewResponse(&pb.UpdateLocationResponse{}), nil
}
