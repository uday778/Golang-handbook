package main

import "fmt"

func main() {
	fmt.Println("if else in Golang")

	loginCount := 23
	var result string
	if loginCount<10 {
		result="regular user"
	} else if loginCount>10{
		result="watch out"
	}else{
		result="exactly 10 login count"
	}
	fmt.Println(result)

	if 9%2 ==0{
		fmt.Println("number is even")
	}else{
		fmt.Println("number is odd")
	}

	if num :=3 ; num < 10 {
		fmt.Println("num is less than 10")
	}else{
		fmt.Println("number is more than 10")
	}
	 //if err != nil {
	//  }



}