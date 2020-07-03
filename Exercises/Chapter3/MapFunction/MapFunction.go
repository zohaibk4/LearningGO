package main

import "fmt"

func mapper(fun func(int) int, arr []int) []int {
	newa := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		newa[i] = fun(arr[i])
	}
	return newa
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{5, 8, 2, 6, 3, 1, 7}
	fun1 := func(z int) int {
		return z * 4
	}
	fun2 := func(z int) int {
		return z * z * z
	}
	fmt.Println(mapper(fun1, arr1))
	fmt.Println(mapper(fun1, arr2))
	fmt.Println(mapper(fun2, arr1))
	fmt.Println(mapper(fun2, arr2))
}
