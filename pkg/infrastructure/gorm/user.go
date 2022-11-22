package gorm

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	domainUser "goarch/pkg/domain/user"
	"goarch/pkg/shared/constants"
	"goarch/pkg/shared/database"
)

type repository struct {
	db *database.Database
}

func UserSetup(database *database.Database) *repository {
	r := &repository{db: database}
	if r.db == nil {
		panic("please provide db")
	}
	return r
}

func (r *repository) Save(entity *domainUser.Entity) (err error) {
	entity.UpdatedAt = time.Now()
	err = r.db.Save(entity).Error
	if err != nil {
		return
	}

	return
}

func (r *repository) FindByEmail(email string) (entity domainUser.Entity, err error) {
	err = r.db.
		Where("email = ?", email).
		First(&entity).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = constants.ErrorUserNotFound
			return
		}
		err = constants.ErrorDatabase
		return
	}

	return
}

func (r *repository) FindById(id int64) (entity domainUser.Entity, err error) {
	err = r.db.
		Where("id = ?", id).
		First(&entity).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = constants.ErrorUserNotFound
			return
		}
		err = constants.ErrorDatabase
		return
	}

	return
}
