package main

import "fmt"

// valoraours operater := not allowed in global scope
//jwttoken:=300000

const LogginToken string= "ggkkjjll" //public that why it have stated with capital letter



func main() {
	var username string = "uday"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n ", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n ", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n ", smallVal)

	var smallfloat float64 = 255.4557714255577
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type: %T \n ", smallfloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n ", anotherVariable)

	//implicit type
	var website="www.google.com"
	fmt.Println(website)

	//no var style 
	//only allowed in main func()
	numberOFUser := 300000.0
	fmt.Println(numberOFUser)

	fmt.Println(LogginToken)
	fmt.Printf("variable  is of type : %T \n", LogginToken)
}
