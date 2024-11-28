package main

import "fmt"

func main() {
	fmt.Println("Welcome to the video of Structs in golang")
	// No Inheritance in Golang ; No Super or parent

	irfan := User{"Irfan", "irfank.ik141@gmail.com", true, 28}
	fmt.Println(irfan)

	fmt.Printf("Irfan details are: %+v\n", irfan)
	fmt.Printf("Name is %v and email is %v.", irfan.Name, irfan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
