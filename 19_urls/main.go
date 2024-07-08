package main

import (
	"fmt"
	"net/url"
)


const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dffgt12"

func main() {
	fmt.Println("welcome to handling URLS in golang")
	fmt.Println(myurl)

	//parsing
	result,_ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("the type of query param are l %T\n",qparams)

	fmt.Println(qparams["coursename"])

	for _,val := range qparams{
		fmt.Println("params is : ",val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}