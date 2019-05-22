package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"com.kaushal/rai/model"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// pgRepository variable available through 'employeeHandler.go'
// var pgRepository *repository.PostgresRepository

const employeeConst = "employee"

func GetPaginatedEmployeeListEmployees_GoChi(w http.ResponseWriter, r *http.Request) {
	reqParms := r.URL.Query()

	start, err := strconv.Atoi(reqParms.Get("start"))
	log.Println("start: ", start)
	if err != nil {
		start = 0
	}

	limit, err2 := strconv.Atoi(reqParms.Get("limit"))
	if err2 != nil {
		limit = 10
	}

	emps, _ := pgRepository.List(start, limit)
	ResponseSuccess(http.StatusOK, w, emps)
}

func CreateEmployee_GoChi(w http.ResponseWriter, r *http.Request) {
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

func GetEmployee_GoChi(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	employee := ctx.Value(employeeConst).(model.Employee)

	ResponseSuccess(http.StatusOK, w, employee)
}

func UpdateEmployee_GoChi(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	log.Println("UpdateEmployee_GoChi :", employee.Id)
	pgRepository.Update(employee)
	ResponseSuccess(http.StatusOK, w, employee)
}

func DeleteEmployee_GoChi(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	employee := ctx.Value("employee").(model.Employee)

	pgRepository.Delete(&employee)
	ResponseSuccess(http.StatusNoContent, w, nil)
}

func EmployeeInContext_GoChi(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		employeeId := chi.URLParam(r, "employeeId")
		employee, err := pgRepository.Get(employeeId)

		if err != nil || employee.Id == "" {
			ResponseError(w, http.StatusNotFound, "The object to be updated could not be found")
			return
		}

		ctx := context.WithValue(r.Context(), employeeConst, employee)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
