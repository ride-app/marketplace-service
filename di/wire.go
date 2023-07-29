//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ride-app/marketplace-service/service"
)

func InitializeService() (*service.EntityServiceServer, error) {
	panic(
		wire.Build(
			statusRepository.NewFirebaseStatusRepository,
			wire.Bind(
				new(statusRepository.StatusRepository),
				new(*statusRepository.FirebaseImpl),
			),
			service.New,
		),
	)
}
