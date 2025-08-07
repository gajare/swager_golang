package main

import (
	"log"
	"net/http"

	_ "github.com/gajare/swager_golang/docs"

	"github.com/gajare/swager_golang/db"
	router "github.com/gajare/swager_golang/routers"
)

// @title Student Microservice API
// @version 1.0
// @description This is a Student CRUD microservice with Swagger docs.
// @host localhost:8080
// @BasePath /
func main() {
	db.InitDB()
	r := router.SetupRouter()
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
