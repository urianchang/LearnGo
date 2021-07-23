package main

import (
	"encoding/json"
	"fmt"
)

var json_string = `{
	"name": "Urian Chang",
	"email": "something@gmail.com",
	"id": 3.14,
	"phone": {
		"home": 99,
		"work": 192
	},
	"male": true,
	"occupation": null
}`

func main() {
	var info map[string]interface{}
	err := json.Unmarshal([]byte(json_string), &info)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(info["name"])
	fmt.Println(info["email"])
	fmt.Println(info["id"])
	fmt.Println(info["phone"].(map[string]interface{})["home"])
	fmt.Println(info["phone"].(map[string]interface{})["work"])
	fmt.Println(info["male"])
	fmt.Println(info["occupation"])
}
