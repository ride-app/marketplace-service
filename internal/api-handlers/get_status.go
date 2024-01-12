package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
)

func (service *MarketplaceServiceServer) GetStatus(ctx context.Context,
	req *connect.Request[pb.GetStatusRequest]) (*connect.Response[pb.GetStatusResponse], error) {
	log := service.logger.WithFields(map[string]string{
		"method": "GetStatus",
	})

	validator, err := protovalidate.New()
	if err != nil {
		log.WithError(err).Info("Failed to initialize validator")

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if err := validator.Validate(req.Msg); err != nil {
		log.WithError(err).Info("Invalid request")

		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.Name, "/")[1]

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

	if status == nil {
		log.Info("status not found")
		return nil, connect.NewError(connect.CodeNotFound, errors.New("status not found"))
	}

	res := connect.NewResponse(&pb.GetStatusResponse{
		Status: status,
	})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("status found")
	return res, nil
}
