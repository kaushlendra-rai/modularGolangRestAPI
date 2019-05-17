package handler

import (
	"github.com/gorilla/mux"
)

func Initialize() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/employee", GetPaginatedEmployeeListEmployees).Methods("GET")
	router.HandleFunc("/employee/{employeeId}", GetEmployee).Methods("GET")
	router.HandleFunc("/employee", CreateEmployee).Methods("POST")
	router.HandleFunc("/employee/{employeeId}", UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{employeeId}", DeleteEmployee).Methods("DELETE")

	return router
}
