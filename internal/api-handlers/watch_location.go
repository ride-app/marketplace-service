package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	locationrepository "github.com/ride-app/marketplace-service/internal/repositories/location"
)

func (service *MarketplaceServiceServer) WatchLocation(ctx context.Context,
	req *connect.Request[pb.WatchLocationRequest], stream *connect.ServerStream[pb.WatchLocationResponse]) error {
	log := service.logger.WithFields(map[string]string{
		"method": "WatchLocation",
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

	uid := strings.Split(req.Msg.Name, "/")[1]

	log.Debug("uid: ", uid)
	log.Debug("request header uid: ", req.Header().Get("uid"))

	if uid != req.Header().Get("uid") {
		return connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	locationResponseStream := make(chan *locationrepository.LocationStreamResponse)

	log.Info("starting Location listener goroutine")
	go func() {
		service.locationRepository.ListenLocation(ctx, log, uid, locationResponseStream)
	}()

	for LocationResponse := range locationResponseStream {
		log.Info("got Location update")
		Location := LocationResponse.Location
		err := LocationResponse.Error

		if err != nil {
			log.WithError(err).Error("failed to get Location")
			return connect.NewError(connect.CodeInternal, err)
		}

		if Location == nil {
			log.Info("location not found")
			return connect.NewError(connect.CodeNotFound, errors.New("location not found"))
		}

		res := &pb.WatchLocationResponse{
			Location: Location,
		}

		if err := validator.Validate(res); err != nil {
			log.WithError(err).Error("Invalid response")
			return connect.NewError(connect.CodeInternal, err)
		}

		log.Info("sending Location update")
		err = stream.Send(res)

		if err != nil {
			log.WithError(err).Error("failed to send response")
			return connect.NewError(connect.CodeInternal, err)
		}
	}

	return nil
}
