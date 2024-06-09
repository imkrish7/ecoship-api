package model

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Buyer struct {
	Id        string    `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar" json:"name"`
	Email     string    `gorm:"type:varchar" json:"email"`
	Password  string    `gorm:"type:varchar" json:"password"`
	Phone     string    `gorm:"type:varchar(12)" json:"phone"`
	Verified  bool      `gorm:"type:boolean" json:"verified"`
	CreateAt  time.Time `gorm:"autoCreateTime:false" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:false" json:"updated_at"`
}

func (base *Buyer) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	time := time.Now().UnixMilli()
	log.Print(time)
	tx.Statement.SetColumn("Id", uuid)
	tx.Statement.SetColumn("CreatedAt", int64(time))
	tx.Statement.SetColumn("UpdatedAt", int64(time))
	return nil
}

func (base *Buyer) BeforeUpdate(tx *gorm.DB) error {
	time := time.Now().UnixMilli()

	tx.Statement.SetColumn("UpdatedAt", int64(time))
	return nil
}
