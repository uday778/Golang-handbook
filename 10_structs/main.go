package main

import "fmt"


type User struct{
		Name string
		Email string
		Status bool
		Age int
	}
func main() {
	fmt.Println("Structs in golang")
	//no inheritance in no super or parent
	
	uday := User{
		Name:   "Uday",
		Email:  "uday@gmail.com",
		Status: true,
		Age:    21,
	}
	fmt.Println(uday)
	fmt.Printf("uday details : %+v\n",uday)
	fmt.Printf("uday is %v and email is %v.",uday.Name, uday.Email)
	


}
