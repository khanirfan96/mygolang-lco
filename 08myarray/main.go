package main

import "fmt"

func main() {
	fmt.Println("elcome to array in GOLang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Pomegranate"
	fruitList[2] = "Mango"
	fruitList[3] = "Banana"

	fmt.Println("FruitList is: ", fruitList)
	fmt.Println("FruitList Length is: ", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "chilli"}
	fmt.Println("Veg List is: ", vegList)
	fmt.Println("Veg List Lenght is: ", len(vegList))
}
