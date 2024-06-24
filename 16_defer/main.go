package main

import "fmt"

func main() {
	defer fmt.Println(" world")
	defer fmt.Println("One")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("welcome to defer in golang")
	myDefer()
	
//defer keyword cutoff and place before the ending curly barcket
}


//   three 1 ,two 2 ,one3 ,world4


// 0,1,2,3,4
//welcome ,43210,three,two,one

func myDefer()  {
	for i:=0;i<5;i++{
		defer fmt.Print(i)
	}
}