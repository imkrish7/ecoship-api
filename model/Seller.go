package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Seller struct {
	Id        string    `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar" json:"name"`
	Email     string    `gorm:"type:varchar" json:"email"`
	Password  string    `gorm:"type:varchar" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	Phone     string    `gorm:"type:varchar(12)" json:"phone"`
	Verified  bool      `gorm:"type:boolean" json:"verified"`
	KycStatus bool      `gorm:"type:boolean" json:"kyc_status"`
	KycId     string    `gorm:"type:uuid" json:"kyc_id"`
	KYC       KYC       `gorm:"foreignKey:KycId"`
}

func (base *Seller) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	time := time.Now().UnixMilli()

	tx.Statement.SetColumn("Id", uuid)
	tx.Statement.SetColumn("CreatedAt", time)
	tx.Statement.SetColumn("UpdatedAt", time)
	return nil
}

func (base *Seller) BeforeUpdate(tx *gorm.DB) error {
	time := time.Now().UnixMilli()

	tx.Statement.SetColumn("UpdatedAt", time)
	return nil
}
