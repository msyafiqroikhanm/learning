package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func Json() {
	SubTitle("Decode JSON To Struct Object")
	var jsonString = `{"Name":"John Doe", "Age":27}`
	var jsonData = []byte(jsonString)

	var data User
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)

	SubTitle("Decode JSON to map")
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])

	SubTitle("Decode Array Json to Array Object")
	var jsonArrayString = `[
		{"Name":"John Name", "Age":23},
		{"Name":"John Wick", "Age":25}
	]`

	var dataArray []User
	err = json.Unmarshal([]byte(jsonArrayString), &dataArray)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1:", dataArray[0].FullName)
	fmt.Println("user 2:", dataArray[1].FullName)

	SubTitle("Encode Object to JSON String")
	var object = []User{{"john doe", 29}, {"doe john", 23}}
	var newJsonData, newErr = json.Marshal(object)
	if newErr != nil {
		fmt.Println(err.Error())
		return
	}
	var newJsonString = string(newJsonData)
	fmt.Println(newJsonString)
}
