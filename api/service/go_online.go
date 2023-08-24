package service

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	pb "github.com/ride-app/marketplace-service/api/gen/ride/marketplace/v1alpha1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (service *MarketplaceServiceServer) GoOnline(ctx context.Context,
	req *connect.Request[pb.GoOnlineRequest]) (*connect.Response[pb.GoOnlineResponse], error) {
	log := service.logger.WithFields(map[string]string{
		"method": "GoOnline",
	})

 	if err := req.Msg.Validate(); err != nil {
   log.WithError(err).Info("Invalid request")
   return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("Invalid request"))
 	}

	uid := strings.Split(req.Msg.Name, "/")[1]

	log.Debug("uid: ", uid)
	log.Debug("Request header uid: ", req.Header().Get("uid"))

 	if uid != req.Header().Get("uid") {
   log.Info("Permission denied")
   return nil, connect.NewError(connect.CodePermissionDenied, errors.New("Permission denied"))
 	}

	driver, err := service.driverRepository.GetDriver(ctx, log, uid, req.Header().Get("Authorization"))

	if err != nil {
		log.WithError(err).Error("Failed to get driver")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

 	if driver == nil {
   log.Info("Driver not found")
   return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("Driver not found"))
 	}

	wallet, err := service.walletrepository.GetWallet(ctx, log, uid, req.Header().Get("Authorization"))

	if err != nil {
		log.WithError(err).Error("Failed to get wallet")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if wallet.Balance <= 0 {
  log.Info("Insufficient wallet balance: ", wallet.Balance)
  return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("Insufficient wallet balance: "+strconv.Itoa(wallet.Balance)))
	}

	vehicle, err := service.driverRepository.GetVehicle(ctx, log, uid, req.Header().Get("Authorization"))

	if err != nil {
		log.WithError(err).Error("Failed to get vehicle")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

 	if vehicle == nil {
   log.Info("Vehicle not found")
   return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("Vehicle not found"))
 	}

	status, err := service.statusRepository.GoOnline(ctx, log, uid, vehicle)

	if err != nil {
		log.WithError(err).Error("Failed to go online")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Status: ", status.Online)

	updateTime, err := service.locationRepository.UpdateLocation(ctx, log, uid, req.Msg.Location)

	if err != nil {
		log.WithError(err).Error("Failed to update location")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	status.UpdateTime = timestamppb.New(*updateTime)

	res := &pb.GoOnlineResponse{
		Status: status,
	}

 	if err := res.Validate(); err != nil {
   log.WithError(err).Error("Failed to update location")
   return nil, connect.NewError(connect.CodeInternal, errors.New("Failed to update location"))
 	}

 log.Info("Driver is online")
	return connect.NewResponse(res), nil
}
