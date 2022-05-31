package database

import (
	"fmt"
	"golang/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	fmt.Println("sentinel1")
	dsn := fmt.Sprintf(

		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	fmt.Println("sentinel2")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println("sentinel3")
	if err != nil {
		fmt.Println("sentinel4")
		fmt.Println(err)
		return db, err
	}

	fmt.Println("sentinel5")
	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Book{})
}
