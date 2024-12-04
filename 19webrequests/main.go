package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	fmt.Println("Wlecome to the video of Web Request in Golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println("Main Content is: ", content)
}
