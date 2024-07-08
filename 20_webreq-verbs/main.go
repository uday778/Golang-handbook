package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to webrequests and verbs in Golang")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "https://api.freeapi.app/api/v1/public/randomjokes"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	
	defer response.Body.Close()
	fmt.Println("status code :",response.StatusCode)
	fmt.Println("content length is : ",response.ContentLength)
	//create a builder
	var responseString strings.Builder
	content,err:=io.ReadAll(response.Body)
	bytecount,_ := responseString.Write(content)

	fmt.Println("bytecount is : ",bytecount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}
