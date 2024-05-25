package initialize

import (
	"fmt"
	"unity/repository/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

/*
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
*/
func ConnectionDB() {
	var err error
	driver := "mysql"
	//os.Getenv("DRIVER")
	host := "127.0.0.1"
	//os.Getenv("HOST")
	user := "kiru"
	//os.Getenv("USER")
	password := "&I%g2o{icSqC"
	//os.Getenv("PASSWORD")
	dbName := "text"
	//os.Getenv("DBNAME")
	dbPor := "3306"
	//os.Getenv("PORT")

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
