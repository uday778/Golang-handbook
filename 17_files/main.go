package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang.")
	content := "this needs to go in a file - learncodeOnline.in"

	file,err:=os.Create("./mylcogofile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	

	length,err :=  io.WriteString(file,content)
	checkNilErr(err)
	
	fmt.Println("length is :",length)
	defer file.Close()
	readfile("./mylcogofile.txt")


}

func checkNilErr(err error)  {
	if  err != nil{
		panic(err)
	}
}


func readfile(filename string)  {
	databyte,err:=os.ReadFile(filename)
	if err != nil{
		panic(err)
	}
	fmt.Println("text data inside the file is \n", string(databyte))
}