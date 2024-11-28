package main

import "fmt"

func main() {
	fmt.Println("Welcome to the video of if else in Golang")

	loginCount := 8
	var result string

	if loginCount < 10 {
		result = "Regular users"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Hey That's it"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println(("Num is not less than 10"))
	}

	// if err != nil {

	// }

}
