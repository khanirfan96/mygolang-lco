package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/khanirfan96/mygolang-lco/mongodb/router"
)

func main() {
	fmt.Println("Welcome to the video of mongodb connection")

	r := router.Router()

	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening at port 4000....")
}
