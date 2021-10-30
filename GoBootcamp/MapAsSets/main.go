package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)

	for in.Scan() {
		word := strings.ToLower(in.Text())
		words[word] = true
	}

	query := "ly"
	if words[query] {
		fmt.Println("The file contains the word ", query)
		return
	}
	fmt.Println("The file does not contain the word ", query)
}
