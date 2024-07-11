package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello mod in golang")
	greeter()
	r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000",r))
}

func greeter()  {
	fmt.Println("hey there mod users")
}

func serveHome(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("<h1>welcome to golang Gogglers</h1>"))
}