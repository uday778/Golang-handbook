package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza app")
	fmt.Println("please rate our pizza 1 or 5")

	reader :=bufio.NewReader(os.Stdin)
	input,_:= reader.ReadString('\n')
	fmt.Println("thanks for rating ,", input)

	numrating,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err !=nil {
		fmt.Println(err)
	} else{
		fmt.Println("added 1 to the ra6ting " , numrating+1)
	}
}


//strconv package 
        //  .ParseInt
		//  .FormatInt
		//  .ParseFloat
		//  .Atoi
		//  .Itoa


// strconv.Atoi(s string) (int, error): Converts a string to an int.

// strconv.Itoa(i int) string: Converts an int to a string.

// strconv.ParseInt(s string, base int, bitSize int) (int64, error): Converts a string to an int64, allowing specification of the numeric base and bit size.



// strconv.FormatInt(i int64, base int) string: Converts an int64 to a string with a specified numeric base.

