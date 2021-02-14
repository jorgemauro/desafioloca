package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type VehicleRepositoryInterface interface {
	FindVehicle(id string) (*Vehicle, error)
}
type Vehicle struct {
	Base            `valid:"required"`
	Placa           string `json:"placa" valid:"notnull"`
	Modelo          string `json:"modelo" valid:"notnull"`
	ValorHora       string `json:"valorhora" valid:"notnull"`
	Combustivel     string `json:"combustivel"`
	LimitePortaMala string `json:"limiteportamala" valid:"notnull"`
	Categoria       string `json:"categoria"`
}

func (vehicle *Vehicle) isValid() error {
	_, err := govalidator.ValidateStruct(vehicle)
	if err != nil {
		return err
	}
	return nil
}

func NewVehicle(placa string, modelo string, valorhora string, combustivel string,
	limiteportamala string, categoria string) (*Vehicle, error) {
	vehicle := Vehicle{
		Placa:           placa,
		Modelo:          modelo,
		ValorHora:       valorhora,
		Combustivel:     combustivel,
		LimitePortaMala: limiteportamala,
		Categoria:       categoria,
	}
	vehicle.ID = uuid.NewV4().String()
	vehicle.CreatedAt = time.Now()

	err := vehicle.isValid()
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}
