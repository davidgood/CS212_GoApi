package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Student Struct to represent the JSON request payload
type Student struct {
	StudentName string         `json:"studentName"`
	StudentID   string         `json:"studentId"`
	StudentType string         `json:"studentType"`
	Grades      map[string]int `json:"grades"`
}

// Function to handle the POST request
func handleStudentData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var student Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	_, err = fmt.Fprintf(w, "Received student data: %+v\n", student)
	if err != nil {
		return
	}

	fmt.Printf("Received student data: %+v\n", student)
}

// Main function to start the web server
func main() {
	http.HandleFunc("/students", handleStudentData)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
