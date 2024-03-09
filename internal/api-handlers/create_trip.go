package apihandlers

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/aidarkhanov/nanoid"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	usecases "github.com/ride-app/marketplace-service/internal/use-cases"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	INITIAL_SEARCH_RADIUS_M = 125
	MAX_SEARCH_RADIUS_M     = 2000
)

func (service *MarketplaceServiceServer) CreateTrip(ctx context.Context,
	req *connect.Request[pb.CreateTripRequest],
) (*connect.Response[pb.CreateTripResponse], error) {
	log := service.logger.WithFields(map[string]string{
		"method": "CreateTrip",
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

	log.Debug("request header uid: ", req.Header().Get("uid"))
	if req.Msg.Trip.Rider.Name != fmt.Sprintf("users/%s", req.Header().Get("uid")) {
		log.Info("permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	trip := req.Msg.Trip

	if trip.Route.Pickup.PolylineString == nil {
		log.Error("pickup polyline string empty")
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("invalid route"))
	}

	tripId := nanoid.New()
	trip.Name = fmt.Sprintf("trips/%s", tripId)
	log.Debug("trip id: ", tripId)

	createTime, err := service.tripRepository.CreateTrip(ctx, log, trip)
	if err != nil {
		log.WithError(err).Error("failed to create trip")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	trip.CreateTime = timestamppb.New(*createTime)
	trip.UpdateTime = timestamppb.New(*createTime)

	go usecases.SearchDriver(ctx, log, INITIAL_SEARCH_RADIUS_M, MAX_SEARCH_RADIUS_M, req.Msg)

	res := connect.NewResponse(&pb.CreateTripResponse{
		Trip: trip,
	})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("trip found")
	return res, nil
}
