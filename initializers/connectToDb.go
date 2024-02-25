package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() *gorm.DB {
	dsn := os.Getenv("DB")
	if dsn == "" {
		panic("DB environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to db: %w", err))
	}

	DB = db
	return db
}
