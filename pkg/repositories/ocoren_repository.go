package repositories

import (
	"github.com/jinzhu/gorm"
)

type OcorenRepository interface {
	//Insert(ocoren *domain.OcorenFile) (domain.OcorenFile, err error)
}

type OcorenRepositoryDb struct {
	Db *gorm.DB
}

//func (repo OcorenRepositoryDb)
