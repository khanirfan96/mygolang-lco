package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the class of slices in GO")

	var fruitList = []string{"Apple", "mango", "banana", "orange"}
	fmt.Printf("TYpe of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "watermleon", "Lemon")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	// fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 867
	highScores[2] = 500
	highScores[3] = 640

	highScores = append(highScores, 555, 666, 221)

	// fmt.Println("High scores of slces", highScores)

	sort.Ints(highScores)
	// fmt.Println("Sorted Highscores", highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove a value from slices based on index

	var courses = []string{"Reactjs", "Javascript", "HTML", "CSS", "Redux"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
