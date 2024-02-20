package main

import (
	"api/controllers"
	"api/middlewares"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middlewares.LogRequestMiddleware)
	r.Use(middlewares.ErrorHandlerMiddleware)
	r.Use(middlewares.CorsMiddleware)

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.RejectNonJsonRequest)
	api.Use(middlewares.JsonResponseMiddleware)

	r.HandleFunc("/", controllers.Home)

	api.HandleFunc("/customers", controllers.HandleGetAllCustomers).Methods("GET")
	api.HandleFunc("/customers/{id}", controllers.HandleGetOneCustomer).Methods("GET")

	api.HandleFunc("/customers", controllers.HandlePostOneCustomer).Methods("POST")

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "7788"
	}

	fmt.Printf("server running in port %v\n", port)
	http.ListenAndServe("0.0.0.0:"+port, r)
}
