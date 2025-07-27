package common

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/victorblv1/brightest_minds/backend/internal/student"
)

func NewRouter(studentHandler *student.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/students", func(r chi.Router) {
		r.Get("/", studentHandler.GetAll)
		r.Get("/{id}", studentHandler.GetByID)
	})

	return r
}
