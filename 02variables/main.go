package main

import "fmt"

const LoginToken string = "knfjsdnjsnf" // When we declare with captal letter it will consider Public

func main() {
	var userName string = "Irfan"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	var smallFloat float64 = 255.78464348749264878287
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	// if type will be string then it will "".
	var anothervariables int
	fmt.Println(anothervariables)
	fmt.Printf("Variable is of type: %T \n", anothervariables)

	// implicit type
	var website = "www.google.com"
	fmt.Println(website)

	// no var style
	numberofUser := 30000 //Walrus operator
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
