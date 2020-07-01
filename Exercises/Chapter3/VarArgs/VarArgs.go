package main

import "fmt"

func printer(numbs ...int) {
	for _, n := range numbs {
		fmt.Println(n)
	}
}

func main() {
	printer(7395, 1638, 603, 5226, 332, 24, 0)
}
