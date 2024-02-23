package apihandlers

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	triprepository "github.com/ride-app/marketplace-service/internal/repositories/trip"
)

func (service *MarketplaceServiceServer) WatchTrip(ctx context.Context, req *connect.Request[pb.WatchStatusRequest], stream *connect.ServerStream[pb.WatchTripResponse]) error {
	log := service.logger.WithFields(map[string]string{
		"method": "GetTrip",
	})

	validator, err := protovalidate.New()
	if err != nil {
		log.WithError(err).Info("Failed to initialize validator")

		return connect.NewError(connect.CodeInternal, err)
	}

	if err := validator.Validate(req.Msg); err != nil {
		log.WithError(err).Info("Invalid request")

		return connect.NewError(connect.CodeInvalidArgument, err)
	}

	tripId := strings.Split(req.Msg.Name, "/")[1]

	log.Debug("trip id: ", tripId)

	trip, err := service.tripRepository.GetTrip(ctx, log, tripId)

	if err != nil {
		log.WithError(err).Error("failed to get trip")
		return connect.NewError(connect.CodeInternal, err)
	}

	if trip == nil {
		log.Info("trip not found")
		return connect.NewError(connect.CodeNotFound, errors.New("trip not found"))
	}

	log.Debug("request header uid: ", req.Header().Get("uid"))
	if trip.Rider.Name != fmt.Sprintf("users/%s", req.Header().Get("uid")) {
		log.Info("permission denied")
		return connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	tripUpdates := make(chan *triprepository.WatchTripResult)

	go service.tripRepository.WatchTrip(ctx, log, tripId, tripUpdates)

	for tripUpdate := range tripUpdates {

		if tripUpdate.Err != nil {
			log.WithError(err).Error("error getting trip update")
			return connect.NewError(connect.CodeInternal, err)
		}

		res := &pb.WatchTripResponse{
			Trip: tripUpdate.Trip,
		}

		if err := validator.Validate(res); err != nil {
			log.WithError(err).Error("invalid response")
			return connect.NewError(connect.CodeInternal, err)
		}

		if err = stream.Send(res); err != nil {
			log.WithError(err).Error("cannot send response")
			return connect.NewError(connect.CodeInternal, err)
		}
	}

	log.Info("watch trip complete")
	return nil
}
