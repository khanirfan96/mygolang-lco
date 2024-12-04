package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=fnskd983783632"

func main() {
	fmt.Println("Welcome to handling URLs in Golang Tutorial")
	fmt.Println(myurl)

	// parsing
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)   // https
	// fmt.Println(result.Host)     // lco.dev:3000
	// fmt.Println(result.Path)     // learn
	// fmt.Println(result.RawQuery) // coursename=reactjs&paymentid=fnskd983783632
	// fmt.Println(result.Port())   // 3000

	qparams := result.Query()
	fmt.Printf("Type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	// fmt.Println(qparams["paymentid"])

	for _, v := range qparams {
		fmt.Println("Params is: ", v)
	}

	partsofURL := &url.URL{
		Scheme:   "https",
		Host:     "google.com",
		Path:     "/gmail",
		RawQuery: "user=irfank.ik",
	}

	anotherURL := partsofURL.String()

	fmt.Println("My new query is: ", anotherURL)
}
