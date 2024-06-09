package db

import (
	"log"

	"github.com/imkrish7/ecoship-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error in connecting DB")
	}

	db.AutoMigrate(&model.Seller{}, &model.Buyer{}, &model.Product{}, &model.KYC{}, &model.Order{})

	return db, nil
}
