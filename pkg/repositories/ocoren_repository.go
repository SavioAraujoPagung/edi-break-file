package repositories

import (
	"fmt"
	"os"

	"github.com/SavioAraujoPagung/edi-break-file/pkg/ocoren"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type OcorenRepository interface {
	Insert(ocoren *ocoren.OccurrencesFile) (*ocoren.OccurrencesFile, error)
}

type OcorenRepositoryDb struct {
	Db *gorm.DB
}

func (repo OcorenRepositoryDb) Insert(ocoren *ocoren.OccurrencesFile) (*ocoren.OccurrencesFile, error) {
	err := repo.Db.Create(ocoren).Error
	if err != nil {

		return nil, err
	}
	return ocoren, nil
}

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}
	dsn := os.Getenv("dsn")

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&ocoren.OccurrenceProceda{})
	return db
}
