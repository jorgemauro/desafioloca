package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	Name      string `gorm:"column:name;type:varchar(255);not null" valid:"notnull"`
	Number    string `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	Cpf       string `json:"cpf" valid:"notnull"`
	Matricula string `json:"matricula"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(number string, Name string, cpf string, matricula string) (*Account, error) {
	account := Account{
		Name:      Name,
		Number:    number,
		Cpf:       cpf,
		Matricula: matricula,
	}
	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}
