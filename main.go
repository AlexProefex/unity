package main

import (
	_ "unity/docs"
	"unity/routes"
)

// @title Api urls for Unity Application
// @version 1.0
// @description  List of all api services
// @host localhost:8080
// @BasePath /api
func main() {

	//token, _ := utils.GenerateToken(20)
	//utils.ExtractTokenID(token)
	routes.Routes()
}
