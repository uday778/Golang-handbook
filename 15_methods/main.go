//copied code from structs

package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
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
	// fmt.Println(uday)
	// fmt.Printf("uday details : %+v\n",uday)
	// fmt.Printf("uday is %v and email is %v.",uday.Name, uday.Email)

	uday.GetStatus()
	uday.newmail()
	fmt.Printf("uday is %v and email is %v.\n", uday.Name, uday.Email)

	//it does not change the actual values of struct thats why pointers exist
}

//methods
func (u User) GetStatus() {
	fmt.Println("Is user is active:", u.Status)
}
func (u User) newmail() {
	u.Email = "test@go.dev"
	fmt.Println("email of this user is : ", u.Email)
}
