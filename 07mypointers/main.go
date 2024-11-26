package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class on pointers")

	// var ptr *int
	// fmt.Println("value of pointer is: ", ptr)

	myNumber := 23

	var ptr = &myNumber                              //refrence mean &(beacause heree we r refrencing to myNumber)
	fmt.Println("Value of actual pointer is ", ptr)  //0xc00000a0e8
	fmt.Println("Value of actual pointer is ", *ptr) //23

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber) //25

	// opeation will be performed on actual value

}
