package service

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	pb "github.com/ride-app/marketplace-service/api/gen/ride/marketplace/v1alpha1"
	statusrepository "github.com/ride-app/marketplace-service/repositories/status"
)

func (service *MarketplaceServiceServer) WatchStatus(ctx context.Context,
	req *connect.Request[pb.WatchStatusRequest], stream *connect.ServerStream[pb.WatchStatusResponse]) error {
	log := service.logger.WithFields(map[string]string{
		"method": "WatchStatus",
	})

	if err := req.Msg.Validate(); err != nil {
		log.WithError(err).Info("Invalid request")
		return connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.Name, "/")[1]

	log.Debug("uid: ", uid)
	log.Debug("request header uid: ", req.Header().Get("uid"))

	if uid != req.Header().Get("uid") {
		return connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	statusResponseStream := make(chan *statusrepository.StatusStreamResponse)

	log.Info("starting status listener goroutine")
	go func() {
		service.statusRepository.ListenStatus(ctx, log, uid, statusResponseStream)
	}()

	for statusResponse := range statusResponseStream {
		log.Info("got status update")
		status := statusResponse.Status
		err := statusResponse.Error

		if err != nil {
			log.WithError(err).Error("failed to get status")
			return connect.NewError(connect.CodeInternal, err)
		}

		if status == nil {
			log.Info("status not found")
			return connect.NewError(connect.CodeNotFound, errors.New("status not found"))
		}

		res := &pb.WatchStatusResponse{
			Status: status,
		}

		if err := res.Validate(); err != nil {
			log.WithError(err).Error("invalid response")
			return connect.NewError(connect.CodeInternal, err)
		}

		log.Info("sending status update")
		err = stream.Send(res)

		if err != nil {
			log.WithError(err).Error("failed to send response")
			return connect.NewError(connect.CodeInternal, err)
		}
	}

	return nil
}
