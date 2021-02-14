package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jorgemauro/desafioloca/serverloca/domain/model"
)

type LocationRepositoryDb struct {
	Db *gorm.DB
}

func (t *LocationRepositoryDb) Register(Location *model.Location) error {
	err := t.Db.Create(Location).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *LocationRepositoryDb) Save(Location *model.Location) error {
	err := t.Db.Save(Location).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *LocationRepositoryDb) Find(id string) (*model.Location, error) {
	var Location model.Location
	t.Db.Preload("AccountFrom.Bank").First(&Location, "id = ?", id)

	if Location.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &Location, nil
}
