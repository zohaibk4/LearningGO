package main

import "fmt"

func plusTwo() func(int) int {
	return func(num int) int {
		return num + 2
	}
}

func main() {
	p2 := plusTwo()
	fmt.Printf("%v\n", p2(496))
	pX := plusX(45)
	fmt.Printf("%v\n", pX(246))
}

func plusX(x int) func(int) int {
	return func(num int) int {
		return num + x
	}
}
