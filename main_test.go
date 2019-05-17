package main_test

import (
	"encoding/json"

	"github.com/gorilla/mux"

	//"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"com.kaushal/rai/handler"
	"com.kaushal/rai/model"
)

// ********** CODE COVERAGE ********
// C:\Users\sinkar\go\gocode\src\rai>go test -cover -coverpkg ./... -coverprofile coverage.out
//
// ** Imp Note: If I do not provide the -coverpkg option, it will give 0% coverage as by default the test coverage is done only for package in which test is run.

// Check the code coverage in browser
// C:\Users\sinkar\go\gocode\src\rai>go tool cover -html coverage.out

var router *mux.Router

func TestMain(m *testing.M) {
	router = handler.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestEmployeesResourceCollection(t *testing.T) {
	createEmployeesInBulk(t)

	req, _ := http.NewRequest("GET", "/employee?start=0&limit=5", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	// Convert response to Object and verify that atleast '5' employees are returned.
	var employees []model.Employee
	json.NewDecoder(response.Body).Decode(&employees)

	if 5 != len(employees) {
		t.Errorf("Expected '5' employees, however got only %d", len(employees))
	}
}

func TestEmployeesResourceCollectionWithDefaultPagination(t *testing.T) {
	createEmployeesInBulk(t)

	req, _ := http.NewRequest("GET", "/employee", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	// Convert response to Object and verify that atleast '5' employees are returned.
	var employees []model.Employee
	json.NewDecoder(response.Body).Decode(&employees)

	if len(employees) <= 0 {
		t.Errorf("Expected some employees, however got only %d", len(employees))
	}
}

func TestGetForNonExistentProduct(t *testing.T) {
	req, _ := http.NewRequest("GET", "/employee/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "The object to be updated could not be found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Product not found'. Got '%s'", m["error"])
	}
}

func TestCreateAndGetEmployee(t *testing.T) {
	employeeUri := createEmployee(t)

	emp := getEmployee(t, employeeUri)

	if emp.Id == "" {
		t.Errorf("employee not found %s", employeeUri)
	}
}

func TestUpdateEmployee(t *testing.T) {
	employeeUri := createEmployee(t)
	emp := getEmployee(t, employeeUri)
	updatedName := "Kaushlendra Rai"
	emp.Name = updatedName
	updateEmployee(t, emp)

	updatedEmp := getEmployee(t, employeeUri)

	if updatedEmp.Name != updatedName {
		t.Errorf("employee name not Updated. Employee's expected updated name= %s. Actual name= %s", updatedName, updatedEmp.Name)
	}
}

func TestDeleteEmployee(t *testing.T) {
	employeeUri := createEmployee(t)
	emp := getEmployee(t, employeeUri)

	deleteEmployee(t, emp.Id, http.StatusNoContent)
}

func TestDeleteEmployeeDoesNotExist(t *testing.T) {
	// Sent non-existent id
	deleteEmployee(t, "ID-DOES-NOT-EXIST", http.StatusNotFound)
}

func createEmployee(t *testing.T) string {
	employee := model.Employee{
		Name: "Sonu", Age: 21, JoiningDate: time.Now(), Department: "BIPRD", PersistenceCommons: model.PersistenceCommons{
			CreationTime: time.Now(), LastModifiedTime: time.Now(), CreatedBy: "sinkar", LastModifiedBy: "sinkar"},
	}

	return createEmployeeWithGivenEmployee(t, employee)
}

func createEmployeeWithGivenEmployee(t *testing.T, employee model.Employee) string {
	emp_json_bytes, _ := json.Marshal(employee)
	req, _ := http.NewRequest("POST", "/employee", strings.NewReader(string(emp_json_bytes[:])))
	response := executeRequest(req)
	location := response.Header().Get("Location")

	checkResponseCode(t, http.StatusCreated, response.Code)

	return location
}

func getEmployee(t *testing.T, employeeUri string) model.Employee {
	req, _ := http.NewRequest("GET", employeeUri, nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var employee model.Employee
	json.NewDecoder(response.Body).Decode(&employee)

	return employee
}

func updateEmployee(t *testing.T, updatedEmployee model.Employee) {
	emp_json_bytes, _ := json.Marshal(updatedEmployee)
	req, _ := http.NewRequest("PUT", "/employee/"+updatedEmployee.Id, strings.NewReader(string(emp_json_bytes[:])))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func deleteEmployee(t *testing.T, employeeId string, expectedStatus int) {
	req, _ := http.NewRequest("DELETE", "/employee/"+employeeId, nil)
	response := executeRequest(req)

	checkResponseCode(t, expectedStatus, response.Code)
}

// Table Driven Tests for Bulk object creation
func createEmployeesInBulk(t *testing.T) {
	//t.Parallel()

	employees := []model.Employee{
		{
			Name: "Sonu", Age: 21, JoiningDate: time.Now(), Department: "BIPRD", PersistenceCommons: model.PersistenceCommons{
				CreationTime: time.Now(), LastModifiedTime: time.Now(), CreatedBy: "sinkar", LastModifiedBy: "sinkar"},
		}, {
			Name: "Sonu2", Age: 21, JoiningDate: time.Now(), Department: "BIPRD", PersistenceCommons: model.PersistenceCommons{
				CreationTime: time.Now(), LastModifiedTime: time.Now(), CreatedBy: "sinkar", LastModifiedBy: "sinkar"},
		}, {
			Name: "Sonu3", Age: 21, JoiningDate: time.Now(), Department: "BIPRD", PersistenceCommons: model.PersistenceCommons{
				CreationTime: time.Now(), LastModifiedTime: time.Now(), CreatedBy: "sinkar", LastModifiedBy: "sinkar"},
		}, {
			Name: "Sonu4", Age: 21, JoiningDate: time.Now(), Department: "BIPRD", PersistenceCommons: model.PersistenceCommons{
				CreationTime: time.Now(), LastModifiedTime: time.Now(), CreatedBy: "sinkar", LastModifiedBy: "sinkar"},
		}, {
			Name: "Sonu5", Age: 21, JoiningDate: time.Now(), Department: "BIPRD", PersistenceCommons: model.PersistenceCommons{
				CreationTime: time.Now(), LastModifiedTime: time.Now(), CreatedBy: "sinkar", LastModifiedBy: "sinkar"},
		}, {
			Name: "Sonu6", Age: 21, JoiningDate: time.Now(), Department: "BIPRD", PersistenceCommons: model.PersistenceCommons{
				CreationTime: time.Now(), LastModifiedTime: time.Now(), CreatedBy: "sinkar", LastModifiedBy: "sinkar"},
		},
	}

	// First create employees ensuring that they are created successfully.
	// Then we will test the /employee GET endpoint to ensure that we get atleast some employees in response
	for _, employee := range employees {
		// The below line ensures that the parallel flows get their own copy of data and do not share data between test runs
		emp := employee
		// Below I am using the 'emp.Name' as the Test name. For specific use-cases, we can have more meaningful names.
		t.Run(emp.Name, func(t *testing.T) {
			t.Parallel() // Create employees in parallel.
			createEmployeeWithGivenEmployee(t, emp)
		})
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
