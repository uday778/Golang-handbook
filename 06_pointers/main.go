package main

import "fmt"

func main() {
	fmt.Println("welcome to class on pointers")
	// var  pointer *string
	// var ptr *int
	// //value of empty pointer is <nill>
	// fmt.Println("value of pointer is : ", ptr) 

	//reffrence means ampersand  symbol
	myNumber :=23
	var ptr =&myNumber
	fmt.Println("value of actual pointer is :  ", ptr)//it give refference  address

	fmt.Println("value of actual pointer is :  ", *ptr)//it gives reffrence value =23
	*ptr = *ptr * 2
	fmt.Println("new value is :", myNumber)

}
