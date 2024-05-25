package main

import (
	"fmt"
	"os"
	"path/filepath"
	"unity/routes"

	"github.com/lpernett/godotenv"
)

// @title Api urls for Unity Application
// @version 1.0
// @description  List of all api services
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = godotenv.Load(filepath.Join(pwd, ".env"))

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	routes.Routes()
}
