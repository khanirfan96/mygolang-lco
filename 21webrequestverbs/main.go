package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// const getUrl string = "http://localhost:8000/get"
// const postUrl string = "http://localhost:8000/post"
const postFormUrl string = "http://localhost:8000/postform"

func main() {
	fmt.Println("Welcome to web Get Video")
	// PerformGetRequest(getUrl)
	// PerformPostJsonRequest(postUrl)
	PerformPostFormRequest(postFormUrl)

}

// func PerformGetRequest(url string) {
// 	response, err := http.Get(url)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer response.Body.Close()

// 	fmt.Println("Status code: ", response.StatusCode)
// 	fmt.Println("Content length is: ", response.ContentLength)

// var responseString strings.Builder
// 	content, _ := ioutil.ReadAll(response.Body)
// byteCount, _ := responseString.Write(content)

// 	fmt.Println(string(content)) // first way to read the data

// fmt.Println("Byte conut is: ", byteCount)
// fmt.Println("Byte conut data is: ", responseString.String()) // second way to read the data

// }

// func PerformPostJsonRequest(url string) {

// fake json payload
// 	requestBody := strings.NewReader(`
// 		{
// 			"coursename": "Let's go with golang",
// 			"price": 0,
// 			"platform": "learnCodeOnlie",
// 			"Heritage": "Taj Mahal"
// 		}
// 	`)

// 	response, err := http.Post(url, "application/json", requestBody)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer response.Body.Close()

// 	content, _ := ioutil.ReadAll(response.Body)
// 	fmt.Println(string(content))

// }

func PerformPostFormRequest(formurl string) {

	// fake formData
	data := url.Values{}
	data.Add("firstname", "Irfan")
	data.Add("lastname", "Khan")
	data.Add("email", "irfank.ik141@gmail.com")

	response, err := http.PostForm(formurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
