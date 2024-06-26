package migrations

import (
	"unity/repository/model"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&model.Categoria{})
	DB.AutoMigrate(&model.Usuarios{})
	DB.AutoMigrate(&model.Locacion{})
	DB.AutoMigrate(&model.UsuarioLocacion{})
	DB.AutoMigrate(&model.Recompensa{})
	DB.AutoMigrate(&model.Premio{})
	DB.AutoMigrate(&model.Canjes{})

	//fmt.Println(err)
}
