package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"corsename"`
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
	return

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
	//generate unique id, string
	//append new course into courses

	k:=rand.New(rand.NewSource(time.Now().UnixNano()))
	
	course.CourseId=strconv.Itoa(rand.Intn(100))
	fmt.Println(k.Intn(100))
	courses =append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}


func updateOneCourse(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id rom req

	params := mux.Vars(r)

	//loop, get id, remove opration, add with my ID 

	for index, course:= range courses{
		if course.CourseId==params["id"]{
			courses=append(courses[:index],courses[index+1:]...  )
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId=params["id"]
			courses=append(courses, course)
			json.NewEncoder(w).Encode(course)
			return 

		}
	}
	//TODO : semd a response whwn id is not found
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func deleteOneCourse(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop,id,remove(index,index+1)
	for index,course:= range courses{
		if course.CourseId ==params["id"] {
			courses=append(courses[:index],courses[index+1:]...)
			break
		}
	}
}