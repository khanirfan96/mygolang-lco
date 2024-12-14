package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var signal = []string{"test"}

var waitg sync.WaitGroup //pointers
var mute sync.Mutex      //pointers

func main() {
	// go greeter("hello")
	// greeter("world")
	websiteList := []string{
		"https://instagram.com",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	for _, web := range websiteList {
		fmt.Println(web, "sjdhshd")
		go getStatusCode(web)
		waitg.Add(1)
	}

	waitg.Wait()
	fmt.Println(signal)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond) // not correct way to handle
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endppoint string) {
	defer waitg.Done()

	res, err := http.Get(endppoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
		log.Fatal(err)
	} else {
		mute.Lock()
		signal = append(signal, endppoint)
		mute.Unlock()

		fmt.Printf("%d status code for %s", res.StatusCode, endppoint)
	}
}
