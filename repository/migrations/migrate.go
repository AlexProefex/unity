package migrations

import (
	"unity/repository/model"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&model.Usuarios{})
}
