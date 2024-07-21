package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang ")
	//creating a channel

	mych := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-mych)
	// mych <- 5

	wg.Add(2)

	//send only
	go func(ch chan<- int, wg *sync.WaitGroup) {

		mych <- 0
		close(mych)

		// mych <- 5
		// mych <- 6

		wg.Done()
	}(mych, wg)


	//Receive only

	go func(ch <-chan int, wg *sync.WaitGroup) {
	
		val, ischannelopen := <-mych

		fmt.Println(val)
		fmt.Println(ischannelopen)

		// fmt.Println(<-mych)//listener

		wg.Done()
	}(mych, wg)
	wg.Wait()

}
