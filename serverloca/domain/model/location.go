package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	LocationPending   string = "pending"
	LocationCompleted string = "completed"
	LocationError     string = "error"
	LocationConfirmed string = "confirmed"
)

type LocationRepositoryInterface interface {
	Register(Location *Location) error
	Save(Location *Location) error
	Find(id string) (*Location, error)
}

type Locations struct {
	Location []Location
}

type Location struct {
	Base              `valid:"required"`
	Account           *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid;" valid:"notnull"`
	Vehicle           *Vehicle `valid:"-"`
	VehicleFromID     string   `gorm:"column:vehicle_from_id;type:uuid;" valid:"notnull"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"-"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (t *Location) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	if t.Status != LocationPending && t.Status != LocationCompleted && t.Status != LocationError {
		return errors.New("invalid status for the Location")
	}

	if err != nil {
		return err
	}
	return nil
}

func (t *Location) Complete() error {
	t.Status = LocationCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Location) Cancel(description string) error {
	t.Status = LocationError
	t.CancelDescription = description
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func NewLocation(account *Account, vehicle *Vehicle, amount float64, description string, id string) (*Location, error) {
	Location := Location{
		Account:       account,
		AccountFromID: account.ID,
		Vehicle:       vehicle,
		VehicleFromID: vehicle.ID,
		Amount:        amount,
		Status:        LocationPending,
		Description:   description,
	}
	if id == "" {
		Location.ID = uuid.NewV4().String()
	} else {
		Location.ID = id
	}
	Location.CreatedAt = time.Now()
	err := Location.isValid()
	if err != nil {
		return nil, err
	}
	return &Location, nil
}
