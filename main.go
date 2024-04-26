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
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	routes.Routes()
}
