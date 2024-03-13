package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (service *MarketplaceServiceServer) GoOnline(ctx context.Context,
	req *connect.Request[pb.GoOnlineRequest],
) (*connect.Response[pb.GoOnlineResponse], error) {
	log := service.logger.WithFields(map[string]string{
		"method": "GoOnline",
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
	log.Debug("Request header uid: ", req.Header().Get("uid"))

	if uid != req.Header().Get("uid") {
		log.Info("Permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	driver, err := service.driverRepository.GetDriver(
		ctx,
		log,
		uid,
		req.Header().Get("Authorization"),
	)
	if err != nil {
		log.WithError(err).Error("Failed to get driver")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if driver == nil {
		log.Info("Driver not found")
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("driver not found"))
	}

	wallet, err := service.walletrepository.GetWallet(
		ctx,
		log,
		uid,
		req.Header().Get("Authorization"),
	)
	if err != nil {
		log.WithError(err).Error("Failed to get wallet")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if wallet.Balance < 0 {
		log.Info("insufficient wallet balance: ", wallet.Balance)

		return nil, connect.NewError(
			connect.CodeFailedPrecondition,
			errors.New("insufficient wallet balance"),
		)
	}

	vehicle, err := service.driverRepository.GetVehicle(
		ctx,
		log,
		uid,
		req.Header().Get("authorization"),
	)
	if err != nil {
		log.WithError(err).Error("failed to get vehicle")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if vehicle == nil {
		log.Info("vehicle not found")
		return nil, connect.NewError(
			connect.CodeFailedPrecondition,
			errors.New("vehicle not found"),
		)
	}

	status, err := service.statusRepository.GoOnline(ctx, log, uid, vehicle)
	if err != nil {
		log.WithError(err).Error("failed to go online")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("status: ", status.Online)

	updateTime, err := service.locationRepository.UpdateLocation(ctx, log, uid, req.Msg.Location)
	if err != nil {
		log.WithError(err).Error("failed to update location")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	status.UpdateTime = timestamppb.New(*updateTime)

	res := connect.NewResponse(&pb.GoOnlineResponse{
		Status: status,
	})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("driver is online")
	return res, nil
}
