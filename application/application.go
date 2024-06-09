package application

import (
	"github.com/imkrish7/ecoship-api/config"
	"gorm.io/gorm"
)

type Application struct {
	DB     *gorm.DB
	CONFIG config.Config
}
