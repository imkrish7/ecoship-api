package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	Id          string         `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar" json:"name"`
	Price       int64          `gorm:"type:bigint" json:"price"`
	Description string         `gorm:"type:text" json:"description"`
	IsAvailable bool           `gorm:"type:boolean" json:"is_available"`
	CreatedAt   time.Time      `gorm:"type:timestamp" json:"created_at"`
	CreatedBy   string         `gorm:"type:uuid" json:"created_by"`
	Seller      Seller         `gorm:"foreignKey:CreatedBy"`
	UpdatedAt   time.Time      `gorm:"type:timestamp" json:"updated_at"`
	Stock       int64          `gorm:"type:bigint" json:"stock"`
	Tags        pq.StringArray `gorm:"type:text[]" json:"tags"`
}

func (base *Product) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	time := time.Now().UnixMilli()

	tx.Statement.SetColumn("Id", uuid)
	tx.Statement.SetColumn("CreatedAt", time)
	tx.Statement.SetColumn("UpdatedAt", time)
	return nil
}

func (base *Product) BeforeUpdate(tx *gorm.DB) error {
	time := time.Now().UnixMilli()

	tx.Statement.SetColumn("UpdatedAt", time)
	return nil
}
