package main

import "fmt"

func main() {
	var s stack
	s.push(12)
	s.push(4)
	s.push(9)
	s.pop()
	s.push(2002)
	fmt.Printf("stack %v\n", s)
}
