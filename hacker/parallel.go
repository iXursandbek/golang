package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	a := make([]int32, n)
	for i := 0; i <= len(a); i++ {
		fmt.Scan(&a[i])
	}
	var coreSoni, limit int32
	fmt.Scan(&coreSoni, &limit)
	fmt.Println(minTime(a, coreSoni, limit))
}

func minTime(files []int32, numCores int32, limit int32) int64 {
	var sum int32
	for i := 0; i < len(files); i++ {
		sum += files[i]
	}
	return int64(sum) + int64(limit)/int64(numCores)

}
