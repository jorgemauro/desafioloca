package model

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Location struct {
	ID            string  `json:"id" validate:"required,uuid4"`
	AccountFromID string  `json:"account_from_id" validate:"required,uuid4"`
	VehicleFromID string  `json:"vehicle_from_id" validate:"required,uuid4"`
	Amount        float64 `json:"amount" validate:"required,numeric"`
	HoursTotal    string  `json:"hours_total" validate:"required"`
	Description   string  `json:"description" validate:"required"`
	Status        string  `json:"status" validate:"-"`
	Error         string  `json:"error"`
}

func (t *Location) isValid() error {
	v := validator.New()
	err := v.Struct(t)
	if err != nil {
		fmt.Errorf("Error during Transaction validation: %s", err.Error())
		return err
	}
	return nil
}

func (t *Location) ParseJson(data []byte) error {
	err := json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	err = t.isValid()
	if err != nil {
		return err
	}

	return nil
}

func (t *Location) ToJson() ([]byte, error) {
	err := t.isValid()
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(t)
	if err != nil {
		return nil, nil
	}

	return result, nil
}

func NewTransaction() *Location {
	return &Location{}
}
