package main

import "fmt"

func average(arr [5]float64) (ave float64) {
	total := 0.0

	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	ave = (total / float64(len(arr)))
	return
}

func main() {
	example := [...]float64{638.21, 2.352, 12.23, 32.8888, 99.32}
	fmt.Print(average(example))
}
