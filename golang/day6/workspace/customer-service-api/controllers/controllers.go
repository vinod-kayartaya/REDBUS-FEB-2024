package controllers

import (
	"api/dao"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "customer service end point here")
}

func HandleGetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorMessage{fmt.Sprint(r)})
		}
	}()

	customers := dao.GetAllCustomers()
	json.NewEncoder(w).Encode(customers)
}

func HandleGetOneCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorMessage{fmt.Sprint(r)})
		}
	}()

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	c := dao.GetOneCustomer(id)

	if c == nil {
		w.WriteHeader(http.StatusNotFound)
		err := ErrorMessage{fmt.Sprintf("No customer found for id %d.", id)}
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(*c)
	}

}
