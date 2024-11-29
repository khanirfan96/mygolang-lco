package main

import "fmt"

func main() {
	fmt.Println("Welcome to function video of Golang")
	greeter()

	result := adder(7, 3)
	fmt.Println("Result is: ", result)

	proResult, myMessage := proAdder(2, 3, 4, 5)
	fmt.Println("pro Result is: ", proResult)
	fmt.Println("pro Message is: ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Namaste from Golang")
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi Irfan for ProAdder for Golang"
}
