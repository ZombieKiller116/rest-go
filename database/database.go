package database

import (
	"fmt"
	"github.com/ZombieKiller116/rest-go/models"
	"gorm.io/gorm/logger"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("DB connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	err = db.AutoMigrate(&models.Fact{})
	if err != nil {
		log.Fatal("Failed to running migrations. \n", err)
	}

	DB = DbInstance{
		Db: db,
	}
}
