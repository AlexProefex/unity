package initialize

import (
	"fmt"
	"log"
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
		log.Fatal("Error loading .env file")
	}

}

func ConnectionDB() {
	var err error
	driver := os.Getenv("DRIVER")
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")
	dbPor := os.Getenv("PORT")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbPor, dbName)
	fmt.Println(url + driver)

	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect with database ", driver)
	} else {
		fmt.Println("connect with database", driver)
	}
	migrations.Migrate(DB)

}
