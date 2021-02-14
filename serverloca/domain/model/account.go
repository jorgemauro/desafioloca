package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type AccountRepositoryInterface interface {
	registerAccount(pixKey *Account) (*Account, error)
	AddUser(user *User) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}
type Account struct {
	Base     `valid:"required"`
	Kind     string `gorm:"column:kind;type:varchar(255);not null" valid:"notnull"`
	Login    string `json:"login" gorm:"type:varchar(20)" valid:"notnull"`
	Password string `json:"password" valid:"notnull"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)

	if account.Kind != "matricula" && account.Kind != "cpf" {
		return errors.New("invalid type of key")
	}
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(login string, password string, kind string) (*Account, error) {
	account := Account{
		Kind:     kind,
		Login:    login,
		Password: password,
	}
	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}
