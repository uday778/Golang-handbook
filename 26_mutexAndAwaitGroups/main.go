package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition in Golang ")
	wg:= &sync.WaitGroup{}
	mut := &sync.Mutex{}
	rwmut:= &sync.RWMutex{}


		rwmut.RLock()
		var score =[]int{0}	
		rwmut.RUnlock()

	
	// wg.Add(1) // use wg.add(3) to work for all three funcs
	wg.Add(3)
	go func (wg *sync.WaitGroup , mut *sync.Mutex){
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg,mut)  
	// wg.Add(1)
	go func (wg *sync.WaitGroup , mut *sync.Mutex){
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()

	}(wg,mut)  
		
	// wg.Add(1)
	go func (wg *sync.WaitGroup, mut *sync.Mutex ,rwmut *sync.RWMutex){
		fmt.Println("Three R")
		rwmut.RLock()
		fmt.Println(score)
		score = append(score, 3)
		rwmut.RUnlock()
		wg.Done()
	}(wg,mut,rwmut)  

	wg.Wait()

	rwmut.RLock()
	fmt.Println(score)
	rwmut.RUnlock()
	
}