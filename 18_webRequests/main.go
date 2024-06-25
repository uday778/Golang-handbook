package main

import (
	"fmt"
	"io"
	"net/http"
)



const url= "https://google.com"

func main() {
	fmt.Println("welcome to web Requests in golang")

	response,err :=http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Tresponse is type of : %T\n",response)
	
	defer response.Body.Close() //callers responsibility to close the connection

	databytes,err:=io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)


}