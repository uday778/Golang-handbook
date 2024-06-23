package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	days :=  []string{"sunday","monday","Tuesday","Wednesday","Thursday","Friday","Saturday"}
	fmt.Println(days)

	// for loop 🔁
	for d :=0;d<len(days);d++ {
		fmt.Println(days[d])
	}

	for i:=range days {
		fmt.Println(days[i])
	}


	for index,day :=range days{
		fmt.Printf("index is %v and value is %v\n",index,day)
	}
	for _,day :=range days{
		fmt.Printf("index is and value is %v\n",day)
	}

	rougueValue := 1
	for rougueValue<10{
		if rougueValue==2 {
			goto lco
		}
		if rougueValue==5 {
			rougueValue++
			continue
		}
		fmt.Println("value is : ",rougueValue)
		rougueValue++
	}
	//goto statements 📚
	lco:
		fmt.Println("Jumping at LearnCodeonline")
}