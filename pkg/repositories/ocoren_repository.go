package repositories

import (
	"log"

	"github.com/SavioAraujoPagung/edi-break-file/pkg/ocoren"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OcorenRepository interface {
	InsertProceda(ocoren *ocoren.OccurrencesFile) (*ocoren.OccurrencesFile, error)
	//	InsertProceda(ocoren *ocoren.OccurrencesFile) (*ocoren.OccurrencesFile, error)
}

type OcorenRepositoryDb struct {
	Db *gorm.DB
}

func (repo *OcorenRepositoryDb) InsertProceda(ocoren *ocoren.OccurrencesFile) (*ocoren.OccurrencesFile, error) {
	err := repo.Db.Create(ocoren).Error
	if err != nil {
		return nil, err
	}
	return ocoren, nil
}

func (repo OcorenRepositoryDb) FindAll(ocoren *ocoren.OccurrencesFile) (*ocoren.OccurrencesFile, error) {
	err := repo.Db.Create(ocoren).Error
	if err != nil {
		return nil, err
	}
	return ocoren, nil
}

func connectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=root dbname=break_file_db_dev port=5412 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	return db
}

func FindOccurrenceCode(ocorenCode *ocoren.OccurrenceCode, code int) {
	db := connectDB()
	db.First(ocorenCode, code)
}

func FindAllOccurrences() []ocoren.OccurrenceCode {
	var occurrences []ocoren.OccurrenceCode
	id := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		id = append(id, i)
	}
	db := connectDB()
	db.Find(&occurrences, id)
	return occurrences
}

func Test(test *ocoren.Test) {
	db := connectDB()
	db.First(test, "10")
	db.Create(test)
}
