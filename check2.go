package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {

	jsonFile, _ := os.Open("/Users/mukeshkhod/go/src/GoProject/newlog.json")
	defer jsonFile.Close()

	fmt.Println("Successfully Opened json file")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var user []map[string]interface{}
	json.Unmarshal([]byte(byteValue), &user)

	for key, result := range user {
		allDetails := result["details"].(map[string]interface{})
		fmt.Println("Reading Value for Key :", key)

		// fmt.Println("name :", result["name"], "- email :", result["email"], "- phone :", result["phone"])
		// fmt.Println("Address :", allDetails["detail1"], allDetails["detail2"])
		// fmt.Println("\n")

		for _, result2 := range allDetails {
			fmt.Println(reflect.TypeOf(result2))
		}

	}

}
