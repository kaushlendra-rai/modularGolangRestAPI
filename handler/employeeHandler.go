package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"com.kaushal/rai/model"
	"com.kaushal/rai/repository"
	"strconv"
	"time"
)

var pgRepository *repository.PostgresRepository

func init() {
	var err error
	pgRepository, err = repository.PostgresConnection()

	if err != nil {
		panic("Could not connect to Database. Terminating application...")
	}
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	json.NewDecoder(r.Body).Decode(&employee)

	id := uuid.New().String()
	employee.Id = id
	employee.CreationTime = time.Now()
	employee.LastModifiedTime = time.Now()

	_ = pgRepository.Add(employee)

	w.Header().Set("Location", "/employee/"+id)
	ResponseSuccess(http.StatusCreated, w, employee)
}

func GetPaginatedEmployeeListEmployees(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	start, err := strconv.Atoi(queryValues.Get("start"))
	if err != nil {
		start = 0
	}

	limit, err2 := strconv.Atoi(queryValues.Get("limit"))
	if err2 != nil {
		limit = 10
	}

	emps, _ := pgRepository.List(start, limit)
	ResponseSuccess(http.StatusOK, w, emps)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeId := vars["employeeId"]

	employee, _ := pgRepository.Get(employeeId)

	if employee.Id != "" {
		ResponseSuccess(http.StatusOK, w, employee)
	} else {
		log.Println("Error: Employee not found: ", employeeId)
		ResponseError(w, http.StatusNotFound, "The object to be updated could not be found")
	}
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)

	existingEmployee, _ := pgRepository.Get(employee.Id)

	// Verify that an employee exists before you try to update it
	if existingEmployee.Id != "" {
		pgRepository.Update(employee)
		ResponseSuccess(http.StatusOK, w, employee)
	} else {
		log.Println("Error: Employee not found: ", employee.Id)
		ResponseError(w, http.StatusNotFound, "The object to be updated could not be found")
	}
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeId := vars["employeeId"]

	existingEmployee, _ := pgRepository.Get(employeeId)

	if existingEmployee.Id != "" {
		pgRepository.Delete(&existingEmployee)
		ResponseSuccess(http.StatusNoContent, w, nil)
	} else {
		log.Println("Error: Employee not found: ", employeeId)
		ResponseError(w, http.StatusNotFound, "The object to be deleted could not be found")
	}
}

func ResponseSuccess(code int, w http.ResponseWriter, body interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
