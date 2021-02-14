package usecase

import (
	"errors"
	"log"

	"github.com/jorgemauro/desafioloca/serverloca/domain/model"
)

type LocationUseCase struct {
	LocationRepository model.LocationRepositoryInterface
	AccountRepository  model.AccountRepositoryInterface
	VehicleRepository  model.VehicleRepositoryInterface
}

func (t *LocationUseCase) Register(accountId string, vehicleId string, amount float64, hoursTotal string, id string) (*model.Location, error) {

	account, err := t.AccountRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}
	vehicle, err := t.VehicleRepository.FindVehicle(vehicleId)
	if err != nil {
		return nil, err
	}
	location, err := model.NewLocation(account, vehicle, amount, id, hoursTotal)
	if err != nil {
		return nil, err
	}
	t.LocationRepository.Save(location)
	if location.Base.ID != "" {
		return location, nil
	}

	return nil, errors.New("unable to process this location")

}

func (t *LocationUseCase) Confirm(locationId string) (*model.Location, error) {
	location, err := t.LocationRepository.Find(locationId)
	if err != nil {
		log.Println("Location not found", locationId)
		return nil, err
	}

	location.Status = model.LocationConfirmed
	err = t.LocationRepository.Save(location)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (t *LocationUseCase) Complete(locationId string) (*model.Location, error) {
	location, err := t.LocationRepository.Find(locationId)
	if err != nil {
		log.Println("Location not found", locationId)
		return nil, err
	}

	location.Status = model.LocationCompleted
	err = t.LocationRepository.Save(location)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (t *LocationUseCase) Error(locationId string) (*model.Location, error) {
	location, err := t.LocationRepository.Find(locationId)
	if err != nil {
		return nil, err
	}

	location.Status = model.LocationError

	err = t.LocationRepository.Save(location)
	if err != nil {
		return nil, err
	}

	return location, nil

}
