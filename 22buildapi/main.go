package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullName"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Understanding building API's in Golang")
	r := mux.NewRouter()

	// Seeding
	courses = append(courses, Course{CourseId: "123", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Darshan", Website: "learnreact.dev"}})
	courses = append(courses, Course{CourseId: "101", CourseName: "GoLang", CoursePrice: 199, Author: &Author{Fullname: "Jacob", Website: "learngo.dev"}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{courseId}", getCourseById).Methods("GET")
	r.HandleFunc("/course", addCourse).Methods("POST")
	r.HandleFunc("/course/{courseId}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{courseId}", deleteCourse).Methods("DELETE")

	// Listen to Port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to implementing API in Go</h1>"))
}

// get all courses route
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get course based on id
func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course by id")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["courseId"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

// add one course
func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add one course")
	w.Header().Set("Content-Type", "application/json")

	// Check if body is not empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Body can not to empty")
	}

	// Handle if JSON is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("JSON data cannot be empty")
		return
	}

	// Check if course with same course name already exists
	for _, coursesData := range courses {
		if course.CourseName == coursesData.CourseName {
			json.NewEncoder(w).Encode("Course with same title already exists")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// Generate unique ID and convert into string
	course.CourseId = strconv.Itoa(rand.IntN(100))

	// Append course into courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

// Update one course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab courseId from request
	params := mux.Vars(r)

	// Check if courseId exists and is not empty
	courseId, err := params["courseId"]
	if !err || courseId == "" {
		// Handle the error case
		json.NewEncoder(w).Encode("CourseId is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Loop over courses, match id and update
	for index, course := range courses {
		if course.CourseId == courseId {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = courseId
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// If no matching course was found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No course found with ID: " + courseId)
}

// Delete one course
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab courseId from request
	params := mux.Vars(r)

	// Check if courseId exists and is not empty
	courseId, err := params["courseId"]
	if !err || courseId == "" {
		// Handle the error case
		json.NewEncoder(w).Encode("CourseId is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Loop over courses, match id and remove
	for index, course := range courses {
		if course.CourseId == courseId {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}

	// If no matching course was found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No course found with ID: " + courseId)
}
