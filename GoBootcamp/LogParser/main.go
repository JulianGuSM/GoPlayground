package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fields := strings.Fields(in.Text())
		//fmt.Printf("domain: %s - visits: %s", fields[0], fields[1])
		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		sum[domain] += visits

	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for domain, visits := range sum {
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
}
