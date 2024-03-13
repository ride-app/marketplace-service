package apihandlers

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
)

// TODO: properly implement
func (service *MarketplaceServiceServer) VerifyTrip(ctx context.Context,
	req *connect.Request[pb.VerifyTripRequest],
) (*connect.Response[pb.VerifyTripResponse], error) {
	log := service.logger.WithFields(map[string]string{
		"method": "VerifyTrip",
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

	tripId := strings.Split(req.Msg.Name, "/")[1]

	log.Debug("trip id: ", tripId)

	trip, err := service.tripRepository.GetTrip(ctx, log, tripId)
	if err != nil {
		log.WithError(err).Error("failed to get trip")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if trip == nil {
		log.Info("trip not found")
		return nil, connect.NewError(connect.CodeNotFound, errors.New("trip not found"))
	}

	log.Debug("request header uid: ", req.Header().Get("uid"))
	if trip.Rider.Name != fmt.Sprintf("users/%s", req.Header().Get("uid")) {
		log.Info("permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	res := connect.NewResponse(&pb.VerifyTripResponse{})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("trip found")
	return res, nil
}
