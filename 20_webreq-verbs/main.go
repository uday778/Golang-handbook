package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	
)

func main() {
	fmt.Println("welcome to webrequests and verbs in Golang")
	// PerformGetRequest()
	PerformPostJsonRequest()
}

// func PerformGetRequest() {
// 	const myurl = "https://api.freeapi.app/api/v1/public/randomjokes"
// 	response, err := http.Get(myurl)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()
// 	fmt.Println("status code :",response.StatusCode)
// 	fmt.Println("content length is : ",response.ContentLength)
// 	//create a builder
// 	var responseString strings.Builder
// 	content,err:=io.ReadAll(response.Body)
// 	bytecount,_ := responseString.Write(content)
// 	fmt.Println("bytecount is : ",bytecount)
// 	fmt.Println(responseString.String())
// 	// fmt.Println(string(content))
// }

func PerformPostJsonRequest() {
	// const myurl = "http://localhost:8000/post"
	// //fake json payload
	// requestBody := strings.NewReader(`
	// {
	//  "coursename":"lets go with golang",
	//  "price":"0",
	//  "platform":"learncodeonline.in
	//  "
	// }
	// `)
	// response,err:=http.Post(myurl,"application/json",requestBody)
	// if err != nil {
	// 	panic(err)	
	// }
	// defer response.Body.Close()
	// io.ReadAll(response.Body)
}

func PerformGPostFormRequest(){
	const myurl = "http:/localhost:8000/postform"

	data := url.Values{}

	data.Add("firstname","uday")
	data.Add("lastname","kumar")
	data.Add("email","uday@gmail.com")


	res,err:= http.PostForm(myurl,data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content ,_:= io.ReadAll((res.Body))
	fmt.Println(string(content))
}
