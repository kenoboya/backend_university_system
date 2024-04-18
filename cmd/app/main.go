package main

import (
	"test-crud/internal/app"
)

const configDir = "../../configs/"

// @title Student Swagger API
// @version 2.0
// @description REST API for test.

// @host localhost:8080
func main() {
	app.Run(configDir)
}
