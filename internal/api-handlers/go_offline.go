package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
)

func (service *MarketplaceServiceServer) GoOffline(ctx context.Context,
	req *connect.Request[pb.GoOfflineRequest],
) (*connect.Response[pb.GoOfflineResponse], error) {
	log := service.logger.WithFields(map[string]string{
		"method": "GoOffline",
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

	log.Info("uid: ", uid)
	log.Debug("request header uid: ", req.Header().Get("uid"))

	if uid != req.Header().Get("uid") {
		log.Info("permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	status, err := service.statusRepository.GoOffline(ctx, log, uid)

	log.Info("status: ", status)

	if err != nil {
		log.Error("failed to go offline")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&pb.GoOfflineResponse{
		Status: status,
	})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("driver went offline")
	return res, nil
}
