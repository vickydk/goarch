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
	dbMaster *database.Database
	dbSlave  *database.Database
}

func UserSetup(dbMaster *database.Database, dbSlave *database.Database) *repository {
	r := &repository{dbMaster: dbMaster, dbSlave: dbSlave}
	if r.dbMaster == nil {
		panic("please provide db master")
	}
	if r.dbSlave == nil {
		panic("please provide db slave")
	}
	return r
}

func (r *repository) Save(entity *domainUser.Entity) (err error) {
	entity.UpdatedAt = time.Now()
	err = r.dbMaster.Save(entity).Error
	if err != nil {
		return
	}

	return
}

func (r *repository) FindByEmail(email string) (entity domainUser.Entity, err error) {
	err = r.dbSlave.
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
	err = r.dbSlave.
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
