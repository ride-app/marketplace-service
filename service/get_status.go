package service

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	pb "github.com/ride-app/marketplace-service/api/gen/ride/marketplace/v1alpha1"
	"github.com/sirupsen/logrus"
)

func (service *MarketplaceServiceServer) GetStatus(ctx context.Context,
	req *connect.Request[pb.GetStatusRequest]) (*connect.Response[pb.GetStatusResponse], error) {

	if err := req.Msg.Validate(); err != nil {
		logrus.Info("Invalid request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.Name, "/")[1]

	logrus.Debug("uid: ", uid)
	logrus.Debug("Request header uid: ", req.Header().Get("uid"))

	if uid != req.Header().Get("uid") {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	status, err := service.statusRepository.GetStatus(ctx, uid)

	if err != nil {
		logrus.WithError(err).Error("Failed to get status")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if status == nil {
		logrus.Info("Status not found")
		return nil, connect.NewError(connect.CodeNotFound, errors.New("status not found"))
	}

	res := &pb.GetStatusResponse{
		Status: status,
	}

	if err := res.Validate(); err != nil {
		logrus.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	logrus.Info("Status found")
	return connect.NewResponse(res), nil
}
