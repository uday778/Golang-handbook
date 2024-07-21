package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup  //pointers



func main() {
	// go greeter("Hello")
	// greeter("world")
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()

	
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}








// package main

// import (
// 	"fmt"
// 	"time"
// )


// func main()  {
// 	fmt.Println("Learning Goroutines...")
// 	 go sayhello()
// 	 go sayhi()


// 	time.Sleep(800*time.Millisecond)
	
// }

// func sayhello()  {
// 	fmt.Println("hello world")
// 	// time.Sleep(2000*time.Millisecond)
// 	fmt.Println("say hello func ended")
// }
// func sayhi()  {
// 	fmt.Println("hi princess")
// 	time.Sleep(1000*time.Millisecond)
// }


