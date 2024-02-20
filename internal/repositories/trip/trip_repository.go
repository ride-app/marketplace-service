package triprepository

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TripRepository interface {
	GetTrip(ctx context.Context, log logger.Logger, id string) (*pb.Trip, error)

	CreateTrip(ctx context.Context, log logger.Logger, trip *pb.Trip) (createTime *time.Time, err error)

	WatchTrip(ctx context.Context, log logger.Logger, id string, watchTripResults chan<- *WatchTripResult)

	UpdateTrip(ctx context.Context, log logger.Logger, id string, trip *pb.Trip) (updateTime *time.Time, err error)
}

type WatchTripResult struct {
	Trip *pb.Trip
	Err  error
}

type FirebaseImpl struct {
	firestore *firestore.Client
	auth      *auth.Client
}

func NewFirebaseTripRepository(log logger.Logger, firebaseApp *firebase.App, auth *auth.Client) (*FirebaseImpl, error) {
	firestore, err := firebaseApp.Firestore(context.Background())

	if err != nil {
		log.WithError(err).Error("error initializing firestore client")
		return nil, err
	}

	if err != nil {
		log.WithError(err).Error("error initializing auth client")
		return nil, err
	}

	log.Info("firebase trip repository initialized")
	return &FirebaseImpl{
		firestore: firestore,
		auth:      auth,
	}, nil
}

func (r *FirebaseImpl) GetTrip(ctx context.Context, log logger.Logger, id string) (*pb.Trip, error) {
	log.Info("querying trip from firestore")
	snap, err := r.firestore.Collection("trips").Doc(id).Get(ctx)

	if err != nil {
		log.WithError(err).Error("error querying trip from firestore")
		return nil, err
	}

	trip := &pb.Trip{}

	jsonBytes, err := json.Marshal(snap.Data())

	if err != nil {
		log.WithError(err).Error("could not marshal trip json into bytes")
		return nil, err
	}

	err = protojson.Unmarshal(jsonBytes, trip)

	if err != nil {
		log.WithError(err).Error("could not unmarshal trip json bytes into Trip")
		return nil, err
	}

	trip.CreateTime = timestamppb.New(snap.CreateTime)
	trip.UpdateTime = timestamppb.New(snap.UpdateTime)

	// trip, err := r.docToTrip(ctx, log, snap)

	// if err != nil {
	// 	log.WithError(err).Error("could not convert firebase document to trip")
	// }

	return trip, nil
}

func (r *FirebaseImpl) WatchTrip(ctx context.Context, log logger.Logger, id string, watchTripResults chan<- *WatchTripResult) {
	iterator := r.firestore.Collection("trips").Doc(id).Snapshots(ctx)
	defer iterator.Stop()

	for {
		snap, err := iterator.Next()
		// DeadlineExceeded will be returned when ctx is cancelled.
		if status.Code(err) == codes.DeadlineExceeded {
			watchTripResults <- &WatchTripResult{
				Trip: nil,
				Err:  err,
			}
			close(watchTripResults)
		}

		if err != nil {
			log.WithError(err).Error("could not get next snapshot")
			watchTripResults <- &WatchTripResult{
				Trip: nil,
				Err:  err,
			}
			close(watchTripResults)
		}

		if !snap.Exists() {
			log.WithError(err).Error("document deleted")
			watchTripResults <- &WatchTripResult{
				Trip: nil,
				Err:  err,
			}
			close(watchTripResults)
		}

		trip := &pb.Trip{}

		jsonBytes, err := json.Marshal(snap.Data())

		if err != nil {
			log.WithError(err).Error("could not marshal trip json into bytes")
			watchTripResults <- &WatchTripResult{
				Trip: nil,
				Err:  err,
			}
			close(watchTripResults)
		}

		err = protojson.Unmarshal(jsonBytes, trip)

		if err != nil {
			log.WithError(err).Error("could not unmarshal trip json bytes into Trip")
			watchTripResults <- &WatchTripResult{
				Trip: nil,
				Err:  err,
			}
			close(watchTripResults)
		}

		trip.CreateTime = timestamppb.New(snap.CreateTime)
		trip.UpdateTime = timestamppb.New(snap.UpdateTime)

		if err != nil {
			log.WithError(err).Error("could not convert snapshot to trip")
			watchTripResults <- &WatchTripResult{
				Trip: nil,
				Err:  err,
			}
			close(watchTripResults)
		}

		watchTripResults <- &WatchTripResult{
			Trip: trip,
			Err:  nil,
		}
	}
}

func (r *FirebaseImpl) CreateTrip(ctx context.Context, log logger.Logger, trip *pb.Trip) (createTime *time.Time, err error) {
	trip.Status = pb.Trip_STATUS_PENDING
	trip.Driver = nil
	trip.Route.WalkToPickup = nil
	trip.Route.WalkToDestination = nil
	trip.CreateTime = nil
	trip.UpdateTime = nil
	trip.StartTime = nil
	trip.EndTime = nil

	jsonData, err := protojson.Marshal(trip)

	if err != nil {
		log.WithError(err).Error("could not marshal trip into json")
		return nil, err
	}

	doc := make(map[string]interface{})

	err = json.Unmarshal(jsonData, &doc)

	if err != nil {
		log.WithError(err).Error("could not unmarshal trip json into map")
		return nil, err
	}

	id := strings.Split(trip.Name, "/")[1]
	writeResult, err := r.firestore.Collection("trips").Doc(id).Set(ctx, doc)

	if err != nil {
		log.WithError(err).Error("could not write trip data to firestore")
		return nil, err
	}

	return &writeResult.UpdateTime, nil
}

func (r *FirebaseImpl) UpdateTrip(ctx context.Context, log logger.Logger, trip *pb.Trip) (updateTime *time.Time, err error) {
	trip.CreateTime = nil
	trip.UpdateTime = nil

	jsonData, err := protojson.Marshal(trip)

	if err != nil {
		log.WithError(err).Error("could not marshal trip into json")
		return nil, err
	}

	doc := make(map[string]interface{})

	err = json.Unmarshal(jsonData, &doc)

	if err != nil {
		log.WithError(err).Error("could not unmarshal trip json into map")
		return nil, err
	}

	id := strings.Split(trip.Name, "/")[1]
	writeResult, err := r.firestore.Collection("trips").Doc(id).Set(ctx, doc)

	if err != nil {
		log.WithError(err).Error("could not write trip data to firestore")
		return nil, err
	}

	return &writeResult.UpdateTime, nil
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
