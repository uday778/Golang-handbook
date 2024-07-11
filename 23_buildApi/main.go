package main

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
	return c.CourseId =="" && c.CourseName ==""
}


func main() {

}
