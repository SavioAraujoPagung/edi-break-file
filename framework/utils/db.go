package utils

/*
import (
	"log"
	"os"

	"github.com/SavioAraujoPagung/edi-break-file/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pg"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dns := os.Getenv("dsn")
	dbType := os.Getenv("dbType")
	db, err := gorm.Open(dbType, dns)
	if err != nil {
		log.Fatalf("Error open connecting to database: %v", err)
		panic(err)
	}

	db.AutoMigrate(&domain.OcorenFile{})
	return db
}
*/
