package main

import "net/http"

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/8/30 9:01 下午
 * @Desc:
 */

func main() {
	http.HandleFunc("/hello", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(write http.ResponseWriter, req *http.Request) {
	_, err := write.Write([]byte("hello!world!"))
	if err != nil {
		return
	}
}
