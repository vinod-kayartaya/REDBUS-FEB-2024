package controllers

import (
	"api/dao"
	"api/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "customer service end point here")
}

func HandleGetAllCustomers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dao.GetAllCustomers())
}

func HandleGetOneCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	c := dao.GetOneCustomer(id)

	if c == nil {
		w.WriteHeader(http.StatusNotFound)
		err := model.ErrorMessage{Message: fmt.Sprintf("No customer found for id %d.", id)}
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(*c)
	}

}

func HandlePostOneCustomer(w http.ResponseWriter, r *http.Request) {
	var cust model.Customer
	json.NewDecoder(r.Body).Decode(&cust)
	cust.Id = dao.AddCustomer(cust)
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(cust)
}
