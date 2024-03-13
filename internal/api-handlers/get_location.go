package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
)

func (service *MarketplaceServiceServer) GetLocation(ctx context.Context,
	req *connect.Request[pb.GetLocationRequest],
) (*connect.Response[pb.GetLocationResponse], error) {
	log := service.logger.WithFields(map[string]string{
		"method": "GetLocation",
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

	location, err := service.locationRepository.GetLocation(ctx, log, uid)
	if err != nil {
		log.WithError(err).Error("failed to get location")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if location == nil {
		log.Info("location not found")
		return nil, connect.NewError(connect.CodeNotFound, errors.New("location not found"))
	}

	res := connect.NewResponse(&pb.GetLocationResponse{
		Location: location,
	})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("location found")
	return res, nil
}
