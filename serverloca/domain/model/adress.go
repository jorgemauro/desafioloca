package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Adress struct {
	Base        `valid:"required"`
	CEP         string  `json:"cep" valid:"notnull"`
	Logradouro  string  `json:"logradouro" valid:"notnull"`
	Cidade      string  `json:"cidade" valid:"notnull"`
	Estado      string  `json:"estado" valid:"notnull"`
	Complemento string  `json:"complemento" valid:"notnull"`
	Account     Account `valid:"-"`
}

func (adress *Adress) isValid() error {
	_, err := govalidator.ValidateStruct(adress)
	if err != nil {
		return err
	}
	return nil
}

func NewAdress(cep string,
	logradouro string,
	cidade string,
	estado string,
	complemento string,
	account Account) (*Adress, error) {
	adresss := Adress{
		CEP:         cep,
		Logradouro:  logradouro,
		Cidade:      cidade,
		Estado:      estado,
		Complemento: complemento,
		Account:     account,
	}

	adresss.ID = uuid.NewV4().String()
	adresss.CreatedAt = time.Now()

	err := adresss.isValid()
	if err != nil {
		return nil, err
	}

	return &adresss, nil
}
