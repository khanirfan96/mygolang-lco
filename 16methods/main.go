package main

import "fmt"

func main() {
	fmt.Println("elcome to video of Method in Golang")
	// No Inheritance in Golang ; No Super or parent

	irfan := User{"Irfan", "irfank.ik141@gmail.com", true, 28}
	fmt.Println(irfan)

	fmt.Printf("Irfan details are: %+v\n", irfan)
	fmt.Printf("Name is %v and email is %v. \n", irfan.Name, irfan.Email)

	irfan.GetStatus()
	irfan.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
