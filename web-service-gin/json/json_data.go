package main

import (
	"encoding/json"
	"fmt"
)

type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

func main() {
	data := []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)

	s := myStruct{
		Name:   "PAlonkas",
		Age:    21,
		Status: true,
		Values: []int{21, 34, 56},
	}

	data1, err := json.Marshal(s)
	data, err = json.MarshalIndent(s, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data1)
	fmt.Println()
	fmt.Printf("%s", data)
	fmt.Println()

	var s1 myStruct
	if err := json.Unmarshal(data, &s1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", s1)
}
