package JsonFetching

import (
	. "GoProject/dbReadCreate"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func FetchFromJson() {

	jsonFile, err := os.Open("/Users/mukeshkhod/go/src/GoProject/newlog.json")
	defer jsonFile.Close()
	CheckErr(err)

	fmt.Println("Successfully Opened json file")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var user []map[string]interface{}
	json.Unmarshal([]byte(byteValue), &user)
	UploadData(user)

}
