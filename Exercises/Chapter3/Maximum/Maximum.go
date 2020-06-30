package main

import "fmt"

func maximum(arr [13]int) (max int) {
	max = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return
}

func main() {
	arr := [...]int{3, 1, 4, 7, 23, 642, 9, 34, 2, 0, 432, 4444, 7}
	max := maximum(arr)
	fmt.Print(max)
}
