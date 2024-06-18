package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)
	reader :=  bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza: ")

	//comma ok syntax|| comma err
	input,_ := reader.ReadString('\n') 
	// _,err := reader.ReadString('\n') 
	fmt.Println("Thanks for Reading, " ,input)
	fmt.Printf("Type of this rating is %T, " ,input)
}