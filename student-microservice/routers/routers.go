package router

import (
	controller "github.com/gajare/swager_golang/controllers"

	_ "github.com/gajare/swager_golang/docs"
	"github.com/gorilla/mux"

	// httpSwagger "github.com/swaggo/http-swagger"

	// _ "github.com/gajare/swager_golang/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/students", controller.CreateStudent).Methods("POST")
	r.HandleFunc("/students", controller.GetStudents).Methods("GET")
	r.HandleFunc("/students/{id}", controller.GetStudent).Methods("GET")
	r.HandleFunc("/students/{id}", controller.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", controller.DeleteStudent).Methods("DELETE")

	// Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
