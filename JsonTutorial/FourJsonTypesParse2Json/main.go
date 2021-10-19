package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/8/26 8:52 下午
 * @Desc: json payload maybe has four types
		  1. normal value(eg:"name":"julian", "age":10, isAuth:true)
		  2. empty value(string "name":"" | array "persons:[]" | object "person":{})
		  3. null value(eg: "name":null)
*/

var jsonData1 = `
{
	"name":"julian",
	"age":10,
	"tags":["handsome", "clever","tall"],
	"friend":{
		"name":"jyq",
		"age": 10
	},
	"isDead":false
}
`

type Firend struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Person struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Tags   []string `json:"tags"`
	Friend Firend   `json:"friend"`
	IsDead bool     `json:"isDead"`
}

func main() {
	gu := new(Person)
	//fmt.Printf("address a struct: %p\n", gu)
	//gu := Person{}
	_ = json.Unmarshal([]byte(jsonData1), gu)
	fmt.Printf("Person gu(输出带字段名): %+v\n", gu)
}
