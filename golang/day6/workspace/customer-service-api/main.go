package main

import (
	"api/controllers"
	"api/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middlewares.LogRequestMiddleware)
	r.Use(middlewares.ErrorHandlerMiddleware)

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.RejectNonJsonRequest)
	api.Use(middlewares.JsonResponseMiddleware)

	r.HandleFunc("/", controllers.Home)

	api.HandleFunc("/customers", controllers.HandleGetAllCustomers).Methods("GET")
	api.HandleFunc("/customers/{id}", controllers.HandleGetOneCustomer).Methods("GET")
	
	api.HandleFunc("/customers", controllers.HandlePostOneCustomer).Methods("POST")

	fmt.Println("server running in port 7788")
	http.ListenAndServe(":7788", r)
}
