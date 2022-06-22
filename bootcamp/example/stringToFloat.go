package main

import (
	"fmt"
	"strconv"
)

func main() {
	f := "3.1415373838"
	if s, err := strconv.ParseFloat(f, 32); err == nil {
		fmt.Println(s)
	}

	if s, err := strconv.ParseFloat(f, 64); err == nil {
		fmt.Println(s)
	}
	l := fmt.Sprintf("%f", 123.34)
	fmt.Println(l)
}
