package main

import "fmt"

func main() {
	example := [...]float64{1.645, 5.357, 7.463, 6.087, 2.537}
	total := 0.0

	for i := 0; i < len(example); i++ {
		total += example[i]
	}
	ave := (total / float64(len(example)))
	fmt.Print(ave)
}
