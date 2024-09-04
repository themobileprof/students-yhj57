package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	dbmain *gorm.DB
)

func initLogger() {
	// Set log level based on environment variable
	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "DEBUG":
		gin.SetMode(gin.DebugMode)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	case "INFO":
		gin.SetMode(gin.ReleaseMode)
		log.SetFlags(log.LstdFlags)
	default:
		gin.SetMode(gin.ReleaseMode)
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags)
	}
}

// GetStudents retrieves a list of all students from the database and returns them as JSON.
func GetStudents(c *gin.Context) {
	log.Println("GET /api/v1/students")

	var students []Student
	dbmain.Find(&students)

	c.IndentedJSON(http.StatusOK, students)
}

// PostStudent adds a new student to the database based on the information provided in the request body and returns the created student as JSON.
func PostStudent(c *gin.Context) {
	log.Println("POST /api/v1/student")

	// Bind student information from Request
	var student Student
	if err := c.BindJSON(&student); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Add student to database
	dbmain.Create(&student)

	// Return student information
	c.IndentedJSON(http.StatusCreated, student)
}

// GetStudent retrieves a specific student from the database based on the provided ID and returns the student as JSON.
func GetStudent(c *gin.Context) {
	log.Printf("GET /api/v1/student/%s", c.Param("id"))

	// Get student ID
	id := c.Param("id")

	// Find student in database
	var student Student
	dbmain.First(&student, id)

	// Return student information
	c.IndentedJSON(http.StatusOK, student)
}

// PutStudent updates the information of a specific student in the database based on the provided ID and the information provided in the request body. It returns the updated student as JSON.
func PutStudent(c *gin.Context) {
	log.Printf("PUT /api/v1/student/%s", c.Param("id"))

	// Get student ID
	id := c.Param("id")

	// Find student in database
	var student Student
	dbmain.First(&student, id)

	// Bind student information from Request
	if err := c.BindJSON(&student); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Save student information
	dbmain.Save(&student)

	// Return student information
	c.JSON(http.StatusOK, student)
}

// DeleteStudent deletes a specific student from the database based on the provided ID and returns the deleted student as JSON.
func DeleteStudent(c *gin.Context) {
	log.Printf("DELETE /api/v1/student/%s", c.Param("id"))

	// Get student ID
	id := c.Param("id")

	// Find student in database
	var student Student
	dbmain.First(&student, id)

	// Delete student from database
	dbmain.Delete(&student)

	// Return student information
	c.IndentedJSON(http.StatusOK, student)
}

// HealthCheck returns a simple health check response.
func HealthCheck(c *gin.Context) {
	// healthcheck the database
	// if err := db.DB().Ping(); err != nil {
	// 	log.Printf("Failed to ping database: %v", err)
	// 	http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func main() {
	initLogger()

	var err error

	dbmain, err = db()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()

	// API routes grouping
	v1 := router.Group("/api/v1")
	{
		v1.GET("/students", GetStudents)
		v1.POST("/student", PostStudent)
		v1.GET("/student/:id", GetStudent)
		v1.PUT("/student/:id", PutStudent)
		v1.DELETE("/student/:id", DeleteStudent)

		v1.GET("/health", HealthCheck)

	}

	router.Run(":8080")
}
