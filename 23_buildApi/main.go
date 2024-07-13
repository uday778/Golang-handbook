package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake Db

var courses []Course

//middleware, helper - file
func(c *Course) IsEmpty() bool {
	// return c.CourseId =="" && c.CourseName ==""
	return  c.CourseName ==""
}

func main()  {
	fmt.Println("API-Learncodeonline.in")
	r:= mux.NewRouter()

	//seeding

	courses= append(courses, Course{CourseId: "2",CourseName: "React js",CoursePrice: 299, Author: &Author{Fullname: "uday kumar",Website: "lco.dev"}})

	courses= append(courses, Course{CourseId: "3",CourseName: "next js", CoursePrice: 399, Author: &Author{Fullname: "rishi",Website: "lco.dev"}})

	courses= append(courses, Course{CourseId: "4",CourseName: "express",CoursePrice: 499, Author: &Author{Fullname: "snake",Website: "go.dev"}})



	//routing

	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")



	//listen to aport
	log.Fatal(http.ListenAndServe(":4000",r))



}



//controllers -- file

//serve home route

func serveHome(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("<h1>welcome to api by chai aur code</h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab it from request
	params := mux.Vars(r)
	
	//loop through courses find matching id and the retutrn the res
	for _,course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	

}

func createOneCourse(w http.ResponseWriter, r * http.Request)  {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")
	//what if : body is empty
	if r.Body ==nil{
		json.NewEncoder(w).Encode("Please send some data")
	}
	//what about empty data -{}
	var course Course
	_=json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
	}

	//todo : check only if title is duplicate
	//loop, title matched with course.coursename, JSON


	//generate unique id, string
	//append new course into courses

	k:=rand.New(rand.NewSource(time.Now().UnixNano()))
	
	course.CourseId=strconv.Itoa(rand.Intn(100))
	fmt.Println(k.Intn(100))
	courses =append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// First - grab id from request
	params := mux.Vars(r)

	// Loop, find ID, remove, add with updated data
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var updatedCourse Course
			err := json.NewDecoder(r.Body).Decode(&updatedCourse)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	// Send response when ID is not found
	http.Error(w, "Course not found", http.StatusNotFound)
}

func deleteOneCourse(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop,id,remove(index,index+1)
	for index,course:= range courses{
		if course.CourseId ==params["id"] {
			courses=append(courses[:index],courses[index+1:]...)
			// todo : send a confirm or demy response 
			break
		}
	}
	
}