package main

import (
	"fmt"
	// "os"
	// "sort"
)

func main() {
	// fmt.Println("welcome to video on Slices")

	// var fruitList = []string{"apple", "lichi", "peach"}
	// fmt.Fprintln(os.Stdout, []any{"tupe of fruitList is %T \n", fruitList}...)

	// fruitList = append(fruitList, "mango", "banana")
	// fmt.Fprintln(os.Stdout, []any{fruitList}...)

	// fruitList = append(fruitList[1:])
	// fruitList = append(fruitList[1:3]) //the last range is not inclusive
	// fmt.Println(fruitList)

	// highscore := make([]int, 4)
	// highscore[0] = 234
	// highscore[1] = 945
	// highscore[2] = 465
	// highscore[3] = 897
	// // highscore[4] = 999 //ourof range or bound

	// highscore = append(highscore, 555, 666, 322)
	// fmt.Println(highscore)

	// fmt.Println(sort.IntsAreSorted(highscore))
	// sort.Ints(highscore)
	// fmt.Println(sort.IntsAreSorted(highscore))
	// fmt.Println(highscore)

	//How to remove a value from slices based on index

	var courses=[]string{"react","java","python","javascript"}
	fmt.Println(courses)

	var index int=2
	courses=append(courses[:index] , courses[index+1])
	fmt.Println(courses)


	//Creatin slice
	// Using a slice literal
   slice := []int{1, 2, 3, 4, 5}

// Using make to create a slice with a specific length and capacity

// slice := make([]int, 5) 
// Creates a slice with length 5 and capacity 5

// var slice []int  // nil slice
// slice = make([]int, 0)


// Slicing Operations

subSlice := slice[1:3]  // subSlice will be [2, 3]
fmt.Println(subSlice)


//Modifying Slice
slice[0] = 10  // Modify an element
slice = append(slice, 4, 5)  // Append elements to slice

//length and capacity
fmt.Println(len(slice))  // prints 5
fmt.Println(cap(slice))  // prints 5

// Iterating over slices
for index, value := range slice {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}



}
