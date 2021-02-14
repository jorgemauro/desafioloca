package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Base      `valid:"required"`
	Name      string   `gorm:"column:name;type:varchar(255);not null" valid:"notnull"`
	Birth     string   `json:"birth" gorm:"type:varchar(20)" valid:"-"`
	Cpf       string   `json:"cpf" gorm:"type:varchar(20)" valid:"-"`
	Matricula string   `json:"matricula" gorm:"type:varchar(20)" valid:"-"`
	Adress    *Adress  `valid:"-"`
	AdressID  string   `gorm:"column:adress_id;type:uuid" valid:"-"`
	Account   *Account `valid:"-"`
	AccountID string   `gorm:"column:account_id;type:uuid" valid:"-"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(birth string, name string, cpf string, matricula string, adress *Adress, account *Account) (*User, error) {
	user := User{
		Name:      name,
		Birth:     birth,
		Cpf:       cpf,
		Matricula: cpf,
		Adress:    adress,
		AdressID:  adress.ID,
		Account:   account,
		AccountID: account.ID,
	}
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
