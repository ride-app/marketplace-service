package triprepository

import (
	"context"
	"time"

	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	types "github.com/ride-app/marketplace-service/internal/utils/types"
)

type TripRepository interface {
	GetTrip(ctx context.Context, log logger.Logger, id string) (*pb.Trip, error)

	CreateTrip(
		ctx context.Context,
		log logger.Logger,
		trip *pb.Trip,
	) (createTime *time.Time, err error)

	WatchTripsCreated(
		ctx context.Context,
		log logger.Logger,
		newTripsResult chan<- *types.StreamResult[*types.Event[*pb.Trip]],
	) error

	WatchTrip(
		ctx context.Context,
		log logger.Logger,
		id string,
		watchTripResults chan<- *WatchTripResult,
	)

	UpdateTrip(
		ctx context.Context,
		log logger.Logger,
		id string,
		trip *pb.Trip,
	) (updateTime *time.Time, err error)
}

// func (r *FirebaseImpl) docToTrip(ctx context.Context, log logger.Logger, doc *firestore.DocumentSnapshot) (*pb.Trip, error) {
// 	log.Debug("Document data: ", doc.Data())
// 	trip := &pb.Trip{}

// 	if !doc.Exists() {
// 		log.Warn("trip not found")
// 		return nil, nil
// 	}

// 	log.Info("fetching rider details")
// 	riderUID := doc.Data()["riderUid"].(string)
// 	rider, err := r.auth.GetUser(ctx, riderUID)

// 	if err != nil {
// 		log.WithError(err).Error("error fetching rider details")
// 		return nil, err
// 	}

// 	trip.Name = fmt.Sprintf("trips/%s", doc.Ref.ID)
// 	trip.Rider = &pb.Trip_Rider{
// 		Name:        fmt.Sprintf("users/%s", rider.UID),
// 		DisplayName: rider.DisplayName,
// 		PhoneNumber: &rider.PhoneNumber,
// 	}
// 	trip.CreateTime = timestamppb.New(doc.CreateTime)
// 	trip.UpdateTime = timestamppb.New(doc.UpdateTime)
// 	trip.Status = pb.Trip_Status(pb.Trip_Status_value[doc.Data()["status"].(string)])
// 	trip.Type = pb.Trip_Type(pb.Trip_Type_value[doc.Data()["type"].(string)])
// 	trip.VehicleType = driverv1alpha1.Vehicle_Type(driverv1alpha1.Vehicle_Type_value[doc.Data()["vehicleType"].(string)])
// 	trip.Passengers = uint32(doc.Data()["passengers"].(int))
// 	trip.PaymentMethod = pb.Trip_PaymentMethod(pb.Trip_PaymentMethod_value[doc.Data()["paymentMethod"].(string)])
// 	trip.Route.Pickup = &pb.Trip_Location{
// 		Coordinates: &latlng.LatLng{
// 			Latitude:  doc.Data()["route.pickup.location.latitude"].(float64),
// 			Longitude: doc.Data()["route.pickup.location.longitude"].(float64),
// 		},
// 		Address: doc.Data()["route.pickup.address"].(string),
// 	}
// 	trip.Route.DropOff = &pb.Trip_Location{
// 		Coordinates: &latlng.LatLng{
// 			Latitude:  doc.Data()["route.dropoff.location.latitude"].(float64),
// 			Longitude: doc.Data()["route.dropoff.location.longitude"].(float64),
// 		},
// 		Address: doc.Data()["route.dropoff.address"].(string),
// 	}

// 	if doc.Data()["route.walk_to_pickup"] != nil {
// 		trip.Route.WalkToPickup = &pb.Trip_Location{
// 			Coordinates: &latlng.LatLng{
// 				Latitude:  doc.Data()["route.walk_to_pickup.location.latitude"].(float64),
// 				Longitude: doc.Data()["route.walk_to_pickup.location.longitude"].(float64),
// 			},
// 			Address: doc.Data()["route.walk_to_pickup.address"].(string),
// 		}
// 	}

// 	if doc.Data()["route.walk_to_pickup"] != nil {
// 		trip.Route.WalkToDestination = &pb.Trip_Location{
// 			Coordinates: &latlng.LatLng{
// 				Latitude:  doc.Data()["route.walk_to_destination.location.latitude"].(float64),
// 				Longitude: doc.Data()["route.walk_to_destination.location.longitude"].(float64),
// 			},
// 			Address: doc.Data()["route.walk_to_destination.address"].(string),
// 		}
// 	}

// 	if doc.Data()["startTime"] != nil {
// 		trip.StartTime = timestamppb.New(doc.Data()["startTime"].(time.Time))
// 	}

// 	if doc.Data()["endTime"] != nil {
// 		trip.EndTime = timestamppb.New(doc.Data()["endTime"].(time.Time))
// 	}

// 	log.Info("checking if trip has a driver assigned")
// 	driverAssigned := doc.Data()["driver.uid"] != nil

// 	if driverAssigned {
// 		log.Info("trip has a driver assigned")
// 		driverUID := doc.Data()["driver.uid"].(string)
// 		driver, err := r.auth.GetUser(ctx, driverUID)

// 		if err != nil {
// 			log.WithError(err).Error("error fetching driver details")
// 			return nil, err
// 		}

// 		trip.Driver = &pb.Trip_Driver{
// 			Name:        fmt.Sprintf("drivers/%s", driver.UID),
// 			DisplayName: driver.DisplayName,
// 			PhoneNumber: &driver.PhoneNumber,
// 		}
// 	}

// 	return trip, nil
// }

// func (r *FirebaseImpl) tripToDoc(ctx context.Context, log logger.Logger, trip *pb.Trip) (*map[string]interface{}, error) {
// 	doc := make(map[string]interface{})

// 	doc["status"] = pb.Trip_Status_name[int32(trip.Status)]
// 	doc["type"] = pb.Trip_Type_name[int32(trip.Type)]
// 	doc["passengers"] = trip.Passengers
// 	doc["riderUid"] = strings.Split(trip.Rider.Name, "/")[1]

// 	return &doc, nil
// }
