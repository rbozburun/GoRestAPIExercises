package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var students = []Student{
	{ID: "1", Name: "John Doe", Class: "1-b", Teachers: "1,2,4"},
}

var classes = []Class{
	{ID: "1", Name: "1-b", Max_size: 20, Student_count: 17},
}

var teachers = []Teacher{
	{ID: "1", Name: "Will", Num_of_Students: 52},
}

type Student struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Class    string `json:"class"`
	Teachers string `json:"teachers"`
}

type Teacher struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Num_of_Students int    `json:"num_of_students"`
}

type Class struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Max_size      int    `json:"max_size"`
	Student_count int    `json:"student_count"`
}

func createStudent(context *gin.Context) {
	var student Student
	err := context.BindJSON(&student)

	// Create student obj with given json value (if there is no err)
	if err == nil && student.ID != "" && student.Name != "" && student.Class != "" && student.Teachers != "" {
		students = append(students, student)
		// Create custom response if created
		context.IndentedJSON(http.StatusCreated, gin.H{"message": "Obj created", "student_id": student.ID})
		return
	} else {
		// Create custom response if not created
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Object cannot created!"})
		return
	}
}

func listStudents(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, students)
}

func getStudentByID(id string) (*Student, error) {
	for i, s := range students {
		if s.ID == id {
			return &students[i], nil
		}
	}

	return nil, errors.New("Student not found!")
}

func getStudent(context *gin.Context) {
	id := context.Param("id")
	fmt.Println("idddd: " + id)

	student, err := getStudentByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student not found!"})
		return
	}

	context.IndentedJSON(http.StatusOK, student)
}

func listClasses(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, classes)
}

func createClass(context *gin.Context) {
	var class Class
	err := context.BindJSON(&class)

	// Create class obj with given json value (if there is no err)
	if err == nil && class.ID != "" && class.Name != "" && class.Max_size != 0 {
		classes = append(classes, class)
		// Create custom response if created
		context.IndentedJSON(http.StatusCreated, gin.H{"message": "Obj created", "class_id": class.ID})
		return
	} else {
		// Create custom response if not created
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Object cannot created!"})
		return
	}
}

func getClassByID(id string) (*Class, error) {
	for i, s := range classes {
		if s.ID == id {
			return &classes[i], nil
		}
	}

	return nil, errors.New("Class not found!")
}

func getClass(context *gin.Context) {
	id := context.Param("id")

	class, err := getClassByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Class not found!"})
		return
	}

	context.IndentedJSON(http.StatusOK, class)
}

func listTeachers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, teachers)
}

func createTeacher(context *gin.Context) {
	var teacher Teacher
	err := context.BindJSON(&teacher)

	// Create teacher obj with given json value (if there is no err)
	if err == nil && teacher.ID != "" && teacher.Name != "" {
		teachers = append(teachers, teacher)
		// Create custom response if created
		context.IndentedJSON(http.StatusCreated, gin.H{"message": "Obj created", "teacher_id": teacher.ID})
		return
	} else {
		// Create custom response if not created
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Object cannot created!"})
		return
	}
}

func getTeacherByID(id string) (*Teacher, error) {
	for i, s := range teachers {
		if s.ID == id {
			return &teachers[i], nil
		}
	}

	return nil, errors.New("Class not found!")
}

func getTeacher(context *gin.Context) {
	id := context.Param("id")

	teacher, err := getTeacherByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Teacher not found!"})
		return
	}

	context.IndentedJSON(http.StatusOK, teacher)
}

func main() {
	router := gin.Default()                 // Router is our server.
	router.GET("/students", listStudents)   // List all student objects
	router.POST("/students", createStudent) // Create new student
	router.GET("/students/:id", getStudent) // Get the specified student with ID

	router.GET("/classes", listClasses)  // List all class objs
	router.POST("/classes", createClass) // Create new class
	router.GET("/classes/:id", getClass) // Get the specified class with ID

	router.GET("/teachers", listTeachers)   // List all teacher objs
	router.POST("/teachers", createTeacher) // Create new teacher
	router.GET("/teachers/:id", getTeacher) // Get the specified teacher with ID

	router.Run("localhost:9090")

}
