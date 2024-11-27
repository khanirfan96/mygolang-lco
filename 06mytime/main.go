package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01/02/2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.October, 01, 14, 14, 0, 0, time.UTC)
	fmt.Println(createdDate)

	fmt.Println("Created format date is: ", createdDate.Format("01/02/2006 Monday"))

}
