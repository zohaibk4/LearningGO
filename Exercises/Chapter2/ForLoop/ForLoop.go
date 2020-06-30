package main

import "fmt"

func main() {
	//Loop 1
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("")
	//Loop 2
	h := 0
Lab:
	if h < 10 {
		fmt.Print(h, " ")
		h++
		goto Lab
	}
	fmt.Println("")
	//Loop 3
	var arr [10]int
	for j := 0; j < 10; j++ {
		arr[j] = j
	}
	fmt.Print(arr)
}
