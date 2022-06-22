package main

import "fmt"

func main() {
	var n int32
	fmt.Scan(&n)
	fizzbuzz(n)
}

func fizzbuzz(n int32) {
	for i := 1; i <= int(n); i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizBuz")
		} else if i%3 == 0 {
			fmt.Println("Fiz")
		} else if i%5 == 0 {
			fmt.Println("Buz")
		} else {
			fmt.Println(i)
		}
	}

}
