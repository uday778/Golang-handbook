package main

import "fmt"

func main() {
	fmt.Println("welcome to arrays in go lang")

	var fruitList [4]string
	fruitList[0]="apple"
	fruitList[1]="orange"
	fruitList[2]="banana"
	fruitList[3]="mango"
	fmt.Println("fruit list is :", fruitList)
	fmt.Println("fruit list is :", len(fruitList))

	var vegList =[5]string{"potato","mirchi","beans"}
	fmt.Println("veg list is : ",vegList)
	fmt.Println("veg list is : ",len(vegList))


}