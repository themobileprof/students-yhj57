package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetStudents(t *testing.T) {
	// Create a test router
	router := gin.Default()
	router.GET("/students", GetStudents)

	// Create a test request
	req, _ := http.NewRequest("GET", "/students", nil)

	// Create a test response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var students []Student
	err := json.Unmarshal(w.Body.Bytes(), &students)
	assert.NoError(t, err)
	assert.NotNil(t, students)
}

func TestPostStudent(t *testing.T) {
	// Create a test router
	router := gin.Default()
	router.POST("/student", PostStudent)

	opt := "someone@gmail.com"

	// Create a test request body
	student := Student{
		FirstName:     "John",
		MiddleName:    nil,
		LastName:      "Doe",
		DoB:           "1990-01-01",
		Gender:        "Male",
		Email:         &opt,
		Address:       "123 Main St",
		GuardianName:  "Jane Doe",
		GuardianPhone: "555-1234",
		GuardianEmail: &opt,
		Grade:         "11",
	}
	body, _ := json.Marshal(student)

	// Create a test request
	req, _ := http.NewRequest("POST", "/student", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a test response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Check the response body
	var createdStudent Student
	err := json.Unmarshal(w.Body.Bytes(), &createdStudent)
	assert.NoError(t, err)
	assert.NotNil(t, createdStudent)
	assert.Equal(t, student.FirstName, createdStudent.FirstName)
	assert.Equal(t, student.Grade, createdStudent.Grade)
}

func TestGetStudent(t *testing.T) {
	// Create a test router
	router := gin.Default()
	router.GET("/student/:id", GetStudent)

	// Create a test request
	req, _ := http.NewRequest("GET", "/student/1", nil)

	// Create a test response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var student Student
	err := json.Unmarshal(w.Body.Bytes(), &student)
	assert.NoError(t, err)
	assert.NotNil(t, student)
}

func TestPutStudent(t *testing.T) {
	// Create a test router
	router := gin.Default()
	router.PUT("/student/:id", PutStudent)

	// Create a test request body
	student := Student{
		FirstName:     "Samuel",
		MiddleName:    nil,
		LastName:      "Johnson",
		DoB:           "1993-01-01",
		Gender:        "Male",
		Email:         nil,
		Address:       "8 Main St",
		GuardianName:  "Jane Doe",
		GuardianPhone: "222-1234",
		GuardianEmail: nil,
		Grade:         "11",
	}
	body, _ := json.Marshal(student)

	// Create a test request
	req, _ := http.NewRequest("PUT", "/student/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a test response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var updatedStudent Student
	err := json.Unmarshal(w.Body.Bytes(), &updatedStudent)
	assert.NoError(t, err)
	assert.NotNil(t, updatedStudent)
	assert.Equal(t, student.FirstName, updatedStudent.FirstName)
	assert.Equal(t, student.Grade, updatedStudent.Grade)
}

func TestDeleteStudent(t *testing.T) {
	// Create a test router
	router := gin.Default()
	router.DELETE("/student/:id", DeleteStudent)

	// Create a test request
	req, _ := http.NewRequest("DELETE", "/student/1", nil)

	// Create a test response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var deletedStudent Student
	err := json.Unmarshal(w.Body.Bytes(), &deletedStudent)
	assert.NoError(t, err)
	assert.NotNil(t, deletedStudent)
}
