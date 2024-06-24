package main

import "fmt"

func main() {
	fmt.Println("welcome to Functions in Go")
	greet()
	result :=adder(3,5)
	fmt.Println("results is : ", result)


	proResults,mymsg:= proadder(2,4,5,6)
	fmt.Println("pro results is : ", proResults)
	fmt.Println("pro message is : ", mymsg)
}
func adder(val1 int,val2 int) int {
	 return val1+val2
} 
// ... the thriple dots are called variatic functions in golang  they can exxpect any values 
func proadder(values ...int)  (int,string){
	total := 0

	for _,val := range values{
		total +=val
	}
	return total, "hi pro results func"
}

func greet()  {
	fmt.Println("hello check check")
}
