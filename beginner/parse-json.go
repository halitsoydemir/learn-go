package beginner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadJsonFile() {
	jsonFile,err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Json read successfully")
	defer jsonFile.Close()
}

func ParseJsonFile() {
	jsonFile,err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue,err := ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValue,&users)
	fmt.Println(users)
	defer jsonFile.Close()
}

func ParseAsUnstructuredJsonFile() {
	jsonFile,err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue,err := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal(byteValue,&result)
	fmt.Println(result)
	defer jsonFile.Close()
}


type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}