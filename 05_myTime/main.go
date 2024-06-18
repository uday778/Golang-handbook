package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study in golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	//format 
	fmt.Println(presentTime.Format("02-01-2006"))
	//if you want day
	fmt.Println(presentTime.Format("02-01-2006 monday"))
	//if you need time 
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 monday"))

	createdDate := time.Date(2020,time.August,12,23,23,0 ,0,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006  monday"))

	ticker := time.NewTicker(1 * time.Second)
go func() {
    for t := range ticker.C {
        fmt.Println("Tick at ", t)
    }
}()
time.Sleep(5 * time.Second)
ticker.Stop()
fmt.Println("Ticker stopped")

}