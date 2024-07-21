package main

import (
	"fmt"
	"git/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MONGODB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000",r))

	fmt.Println("Listening At Port 4000 ...")
	
}