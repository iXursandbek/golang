package main

import "fmt"

func main() {
	s := fmt.Sprintf("Binary: %b\\%b", 4, 5)
	fmt.Println(s)

	a := 2
	fmt.Printf("%#v\n", a)
}
