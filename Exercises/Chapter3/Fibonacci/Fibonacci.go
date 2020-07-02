package main

import "fmt"

func fib(num int) []int {
	var arr = make([]int, num)
	arr[0] = 1
	arr[1] = 1
	for i := 2; i < num; i++ {
		arr[i] = arr[i-2] + arr[i-1]
	}
	return arr
}

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
	fmt.Println(fib(10))
	fmt.Println(fib(20))
}
