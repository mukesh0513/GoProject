package JsonFetching

import (
	. "GoProject/dbReadCreate"
	. "GoProject/jsonStruct"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func FetchFromJson() {

	var user []Users
	jsonFile, err := os.Open("log.json")
	CheckErr(err)

	fmt.Println("Successfully Opened users.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &user)
	UploadData(user)
	defer jsonFile.Close()
}
