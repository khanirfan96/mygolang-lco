package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("elcome to our Pizza App")
	fmt.Println("Please rate between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	// Comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of reading is %T", input)
}
