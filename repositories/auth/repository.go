package repositories

import (
	"context"
	"log"

	"github.com/imkrish7/ecoship-api/model"
	services "github.com/imkrish7/ecoship-api/services/hash"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) CreateUser(ctx context.Context, name string, email string, userType string, password string, phone string) (*gorm.Tx, error) {
	var buyer = model.Buyer{Email: email}
	var seller = model.Seller{Email: email}
	buyerExist := a.db.Find(&buyer)
	sellerExist := a.db.Find(&seller)
	log.Print(buyerExist.Error)
	log.Print(*sellerExist)

	buyer.Name = email
	buyer.Phone = phone
	passwordHash, err := services.HashPassword(password)
	if err != nil {
		log.Fatalf("password hash problem", err)
	}

	buyer.Password = passwordHash

	newBuyer := a.db.Create(&buyer)
	log.Print(newBuyer.Error)
	log.Print(newBuyer.RowsAffected)
	return nil, nil
}
