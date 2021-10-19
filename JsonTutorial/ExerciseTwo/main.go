package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/8/26 8:11 下午
 * @Desc:
 */

type MyStruct struct {
	FieldA string `json:"field_a,omitempty"`
	FieldC string `json:"field_c"`
}

func main() {
	mystruct := MyStruct{
		FieldA: "",
		FieldC: "",
	}

	jsonStr, _ := json.Marshal(mystruct)
	fmt.Printf("json string is : %s\n", string(jsonStr))
}
