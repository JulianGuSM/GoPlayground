package main

import (
	"bufio"
	"fmt"
	"os"
)

// scan an input stream line by line into a buffer
func main() {
	in := bufio.NewScanner(os.Stdin)

	var lines = 0
	for in.Scan() {
		lines++
		fmt.Println("Scanned Text:", in.Text())
	}
	fmt.Printf("There are %d line(s)", lines)

	//fmt.Println("Scanned Text:", in.Text())
	//fmt.Println("Scanned Bytes: ", in.Bytes())

}
