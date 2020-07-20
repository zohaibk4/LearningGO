package main

import "fmt"

func loop1() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
}

func loop2() {
	h := 0
Lab:
	if h < 10 {
		fmt.Print(h, " ")
		h++
		goto Lab
	}
}

func loop3() {
	var arr [10]int
	for j := 0; j < 10; j++ {
		arr[j] = j
	}
	fmt.Print(arr)
}

func main() {
	loop1()
	fmt.Println("")
	loop2()
	fmt.Println("")
	loop3()
}
