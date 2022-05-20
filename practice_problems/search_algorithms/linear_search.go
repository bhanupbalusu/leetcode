package main

import "fmt"

func main() {
	fmt.Println(lsearch([]int32{10, 13, 15, 16, 17, 19}, 15))
}

func lsearch(a1 []int32, n int32) int {
	for i := range a1 {
		if a1[i] == n {
			return i
		}
	}
	return -1
}
