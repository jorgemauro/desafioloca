package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base     `valid:"required"`
	Kind     string `gorm:"column:kind;type:varchar(255);not null" valid:"notnull"`
	Login    string `json:"login" gorm:"type:varchar(20)" valid:"notnull"`
	Password string `json:"password" valid:"notnull"`
	User     *User  `valid:"-"`
	UserID   string `gorm:"column:user_id;type:uuid;not null" valid:"-"`
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

func NewAccount(user *User, login string, password string, kind string) (*Account, error) {
	account := Account{
		Kind:     kind,
		Login:    login,
		Password: password,
		User:     user,
		UserID:   user.ID,
	}
	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}
