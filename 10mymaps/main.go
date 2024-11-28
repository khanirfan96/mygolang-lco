package main

import "fmt"

func main() {
	fmt.Println("Welcome to the maps video")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["react"] = "ReactJs"
	languages["rx"] = "Redux"

	fmt.Println("list of languages: ", languages)
	fmt.Println("js short for: ", languages["js"])

	delete(languages, "rx")
	fmt.Println("deleted languages: ", languages)

	// loops are interesting in go
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v", key, value)
	}
}
