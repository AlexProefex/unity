package config

/*
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	driver := "mysql"   //os.Getenv("DRIVER")
	host := "localhost" //os.Getenv("HOST")
	user := "test"      //os.Getenv("USER")
	password := "test"  //os.Getenv("PASSWORD")
	dbName := "text"    //os.Getenv("DBNAME")
	dbPor := "3600"     //os.Getenv("PORT")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbPor, dbName)
	fmt.Println(url)

	DB, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect with database ", driver)
	} else {
		fmt.Println("connect with database", driver)
	}

	fmt.Println(DB)

}
*/
