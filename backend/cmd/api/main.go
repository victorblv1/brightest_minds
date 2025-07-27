package main

import (
	"log"
	"net/http"

	"github.com/victorblv1/brightest_minds/backend/internal/common"
	"github.com/victorblv1/brightest_minds/backend/internal/student"
)

func main() {
	// Initialize services
	studentRepo := student.NewInMemoryRepository()
	studentService := student.NewService(studentRepo)
	studentHandler := student.NewHandler(studentService)

	// Setup router
	r := common.NewRouter(studentHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
