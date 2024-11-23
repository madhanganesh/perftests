package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	start := time.Now()

	for e := 0; e < 30; e++ {
		var x [1000000]int
		for i := 0; i < 1000000; i++ {
			x[i] = i
		}

		var y [1000000 - 1]int
		for i := 0; i < 1000000-1; i++ {
			y[i] = x[i] + x[i+1]
		}

		sum = 0
		for i := 0; i < 1000000; i += 100 {
			sum += y[i]
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("Elapsed %s\n", elapsed)
	fmt.Println(sum)
}
