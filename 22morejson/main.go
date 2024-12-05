package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // this "-" will hide this key from json
	Tags     []string `json:"tags,omitempty"` // omitempty say if it's null then dont't throw this key
}

func main() {
	fmt.Println("Welcome! to video of more about json in Golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS", 300, "Google.com", "abc123", []string{"Web-dev", "js"}},
		{"Javascript", 500, "Namastedev.com", "abc455", []string{"Backend", "js"}},
		{"Java", 400, "Telusko.in", "Git123", nil},
	}

	//Package this data as json data
	//finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDatafromWeb := []byte(`
        {
                "coursename": "ReactJS",
                "Price": 300,
                "website": "Google.com",
                "tags": ["Web-dev", "js"]
        }	
	`)

	var maincourse course

	result := json.Valid(jsonDatafromWeb)

	if result {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDatafromWeb, &maincourse)
		fmt.Printf("%#v\n", maincourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// Some cases where you just want to add data to key value pair
	var myData map[string]interface{}
	json.Unmarshal(jsonDatafromWeb, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
