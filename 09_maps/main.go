package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("maps in go lang")
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"

	fmt.Println("list of all loanguages ", languages)
	fmt.Println("Js shorts for:  ", languages["js"])
	delete(languages, "rb")
	fmt.Println("list of all loanguages ", languages)

	//loops in maps 
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
	for _, value := range languages {
		fmt.Printf("For key v, value is %v \n", value)
	}

	//Creating and initializing a map
	fruits := map[string]string{
		"apple":  "1",
		"banana": "2",
	}

	//adding an element
	fruits["cherry"] = "3"

	//updating an element
	fruits["apple"] = "10"

	//Accessing elements
	fmt.Println("apple:", fruits["apple"])
	fmt.Println("banana:", fruits["banana"])

	//check if a key exists
	value, ok := fruits["apple"]
	if ok {
		fmt.Println("Apple exists with value:", value)
	} else {
		fmt.Println("apple does no exist")
	}

	//Deleting an element
	delete(fruits, "banana")

	// Iterating over a map
	for key, value := range fruits {
		fmt.Fprintf(os.Stdout, "Key: %s, Value: %d\n", []any{key, value}...)
	}


	



}
