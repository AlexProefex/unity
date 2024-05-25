package initialize

import (
	"fmt"
	"os"
	"path/filepath"
	"unity/repository/migrations"

	"github.com/lpernett/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = godotenv.Load(filepath.Join(pwd, ".env"))

	if err != nil {
		fmt.Println("Error loading .env file")
	}

}

func ConnectionDB() {
	var err error
	//driver := "mysql"
	driver := os.Getenv("DRIVER")
	//host := "127.0.0.1"
	host := os.Getenv("HOST")
	//user := "kiru"
	user := os.Getenv("USUARIO")
	//password := "&I%g2o{icSqC"
	password := os.Getenv("PASSWORD")
	//dbName := "text"
	dbName := os.Getenv("DBNAME")
	//dbPort := "3306"
	dbPort := os.Getenv("PORT")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbPort, dbName)
	fmt.Println(url + driver)

	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect with database ", driver)
	} else {
		fmt.Println("connect with database", driver)
	}
	migrations.Migrate(DB)

}
