package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/gorilla/mux"
)

func InitializeGorillaMuxRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/employee", GetPaginatedEmployeeListEmployees).Methods("GET")
	router.HandleFunc("/employee/{employeeId}", GetEmployee).Methods("GET")
	router.HandleFunc("/employee", CreateEmployee).Methods("POST")
	router.HandleFunc("/employee/{employeeId}", UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{employeeId}", DeleteEmployee).Methods("DELETE")

	return router
}

func InitializeGoChiRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/employee", func(r chi.Router) {
		r.Get("/", GetPaginatedEmployeeListEmployees_GoChi)
		r.Post("/", CreateEmployee_GoChi)

		r.Route("/{employeeId}", func(r chi.Router) {
			r.Use(EmployeeInContext_GoChi)
			r.Get("/", GetEmployee_GoChi)
			r.Put("/", UpdateEmployee_GoChi)
			r.Delete("/", DeleteEmployee_GoChi)
		})
	})

	// Accessible at URL: http://localhost/admin
	r.Mount("/cancelable", cancelableRouter())
	return r
}

func cancelableRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", CancelableRestAPI)

	return r
}
