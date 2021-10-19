package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/8/26 7:25 下午
 * @Desc:
 */

type App struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

var jsonData = `{
        "id": "k34rAT4",
        "title": "My Awesome App"
    }`

func main() {
	var app App
	// []byte -> struct using json.Unmarshal
	err := json.Unmarshal([]byte(jsonData), &app)
	if err != nil {
		fmt.Printf("catch err: %v\n", err)
	}

	fmt.Printf("var app looks like(值的默认格式表示): %v\n", app)
	fmt.Printf("var app looks like(类似%%v，但输出结构体时会添加字段名): %+v\n", app)
	fmt.Printf("var app looks like(值的Go语法表示): %#v\n", app)

	// ignore err is not good in work,just for playground
	// struct -> []byte using json.Marshal
	data, _ := json.Marshal(app)
	fmt.Printf("json data looks like: %s\n", string(data))
}
