package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := make(map[string]interface{})
	data["name"] = "Urian"
	data["occupation"] = nil
	data["phone"] = map[string]interface{}{
		"home": 123,
		"work": 333,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Stringify JSON for printing
	jsonStr := string(jsonData)
	fmt.Println("JSON string:")
	fmt.Println(jsonStr)
}
