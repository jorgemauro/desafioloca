package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/jorgemauro/desafioloca/serverloca/infrastructure/repository"
	"github.com/jorgemauro/desafioloca/serverloca/serverloca/application/usecase"
	"github.com/jorgemauro/desafioloca/serverloca/serverloca/infrastructure/repository"
)

func LocationUseCaseFactory(database *gorm.DB) usecase.LocationUseCase {
	locationRepository := repository.LocationRepositoryDb{Db: database}
	ServerLocaRepository := repository.AccountRepositoryDb{Db: database}
	locationUseCase := usecase.LocationUseCase{
		LocationRepository:   &locationRepository,
		ServerLocaRepository: ServerLocaRepository,
	}

	return locationUseCase
}
