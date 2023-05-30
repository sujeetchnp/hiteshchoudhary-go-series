package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup // pointers
var mut sync.Mutex    // pointers

func main() {
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	for _, webs := range websiteList {
		go getStatusCode(webs)
		wg.Add(1) // keep adding go routines
	}

	wg.Wait() // Wait for go routines to complete
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done() // Once completed it sends Done
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}

/**
Example 1:

func main() {
	go greeter("Hello")
	greeter("World")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
**/
