package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KYC struct {
	Id        string    `gorm:"type:uuid;primaryKey" json:"id"`
	FullName  string    `gorm:"type:varchar" json:"fullname"`
	Age       int64     `gorm:"type:varchar" json:"age"`
	BirthDate string    `gorm:"type:varchar" json:"birth_date"`
	Aadhar    string    `gorm:"type:varchar" json:"aadhar"`
	Pan       string    `gorm:"type:varchar" json:"pan"`
	Address   string    `gorm:"type:varchar" json:"address"`
	Phone     string    `gorm:"type:varchar(12)" json:"phone"`
	CreateAt  time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (base *KYC) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	time := time.Now().UnixMilli()

	tx.Statement.SetColumn("Id", uuid)
	tx.Statement.SetColumn("CreatedAt", time)
	tx.Statement.SetColumn("UpdatedAt", time)
	return nil
}

func (base *KYC) BeforeUpdate(tx *gorm.DB) error {
	time := time.Now().UnixMilli()

	tx.Statement.SetColumn("UpdatedAt", time)
	return nil
}
