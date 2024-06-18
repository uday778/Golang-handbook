package main

import "fmt"

func main() {
	fmt.Println("maps in go lang")
	languages:=make(map[string]string)
	languages["js"]="javascript"
	languages["py"]="python"
	languages["rb"]="ruby"

	fmt.Println("list of all loanguages ", languages)
	fmt.Println("Js shorts for:  ", languages["js"])
	delete(languages,"rb")
	fmt.Println("list of all loanguages ", languages)


	//loops in maps are intresting
	for key,value:= range languages{
		fmt.Printf("For key %v, value is %v \n", key,value)
	}
	for _,value:= range languages{
		fmt.Printf("For key v, value is %v \n", value)
	}
}