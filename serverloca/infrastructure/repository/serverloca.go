package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jorgemauro/desafioloca/serverloca/domain/model"
)

type AccountRepositoryDb struct {
	Db *gorm.DB
}

func (r AccountRepositoryDb) AddAdress(adress *model.Adress) error {
	err := r.Db.Create(adress).Error
	if err != nil {
		return err
	}
	return nil
}

func (r AccountRepositoryDb) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

func (r AccountRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account
	r.Db.Preload("account").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account found")
	}
	return &account, nil
}

func (r AccountRepositoryDb) FindAdress(id string) (*model.Adress, error) {
	var adress model.Adress
	r.Db.First(&adress, "id = ?", id)

	if adress.ID == "" {
		return nil, fmt.Errorf("no bank found")
	}
	return &adress, nil
}
