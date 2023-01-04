package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user USER

	json.NewDecoder(r.Body).Decode(&user)

	Database.Create(&user)
	json.NewEncoder(w).Encode(user)

}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []USER

	Database.Find(&users)

	json.NewEncoder(w).Encode(users)

}
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user USER

	Database.First(&user, mux.Vars(r)["id"])

	json.NewEncoder(w).Encode(user)

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user USER

	Database.First(&user, mux.Vars(r)["id"])

	json.NewDecoder(r.Body).Decode(&user)

	Database.Save(&user)

	json.NewEncoder(w).Encode(user)

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user USER

	Database.Delete(&user, mux.Vars(r)["id"])

	Database.Save(&user)

	json.NewEncoder(w).Encode("Employee Deleted Successfully")

}
