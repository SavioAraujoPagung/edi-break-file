package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=Sistemadetrocadeitens password=1234 dbname=break_file port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Printf("db: %v\n", db)
	if err != nil {
		log.Panic(err)
	}

}
