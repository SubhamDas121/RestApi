package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()

	r.HandleFunc("/user", CreateEmployee).Methods("POST")
	r.HandleFunc("/users", getEmployees).Methods("GET")
	r.HandleFunc("/user/{id}", GetEmployeeByID).Methods("GET")
	r.HandleFunc("/user/{id}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/user/{id}", DeleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}
