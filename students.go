package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStudents retrieves a list of all students from the database and returns them as JSON.
func GetStudents(c *gin.Context) {
	db, err := db()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var students []Student
	db.Find(&students)

	c.IndentedJSON(http.StatusOK, students)
}

// PostStudent adds a new student to the database based on the information provided in the request body and returns the created student as JSON.
func PostStudent(c *gin.Context) {
	db, err := db()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Bind student information from Request
	var student Student
	if err := c.BindJSON(&student); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Add student to database
	db.Create(&student)

	// Return student information
	c.IndentedJSON(http.StatusCreated, student)
}

// GetStudent retrieves a specific student from the database based on the provided ID and returns the student as JSON.
func GetStudent(c *gin.Context) {
	db, err := db()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get student ID
	id := c.Param("id")

	// Find student in database
	var student Student
	db.First(&student, id)

	// Return student information
	c.IndentedJSON(http.StatusOK, student)
}

// PutStudent updates the information of a specific student in the database based on the provided ID and the information provided in the request body. It returns the updated student as JSON.
func PutStudent(c *gin.Context) {
	db, err := db()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get student ID
	id := c.Param("id")

	// Find student in database
	var student Student
	db.First(&student, id)

	// Bind student information from Request
	if err := c.BindJSON(&student); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Save student information
	db.Save(&student)

	// Return student information
	c.IndentedJSON(http.StatusOK, student)
}

// DeleteStudent deletes a specific student from the database based on the provided ID and returns the deleted student as JSON.
func DeleteStudent(c *gin.Context) {
	db, err := db()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get student ID
	id := c.Param("id")

	// Find student in database
	var student Student
	db.First(&student, id)

	// Delete student from database
	db.Delete(&student)

	// Return student information
	c.IndentedJSON(http.StatusOK, student)
}

func main() {
	router := gin.Default()

	router.GET("/students", GetStudents)
	router.POST("/student", func(c *gin.Context) {
		PostStudent(c)
	})
	router.GET("/student/:id", GetStudent)
	router.PUT("/student/:id", func(c *gin.Context) {
		PutStudent(c)
	})
	router.DELETE("/student/:id", DeleteStudent)

	router.Run(":8080")
}
