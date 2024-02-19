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
	r.Use(middlewares.RejectNonJsonRequest)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/customers", controllers.HandleGetAllCustomers).Methods("GET")
	r.HandleFunc("/api/customers/{id}", controllers.HandleGetOneCustomer).Methods("GET")

	fmt.Println("server running in port 7788")
	http.ListenAndServe(":7788", r)
}
