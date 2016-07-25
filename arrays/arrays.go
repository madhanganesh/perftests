package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	start := time.Now()

	for e := 0; e < 30; e++ {
		x := make([]int, 0, 1000000)
		for i := 0; i < 1000000; i++ {
			x = append(x, i)
		}

		y := make([]int, 0, 1000000-1)
		for i := 0; i < 1000000-1; i++ {
			y = append(y, x[i]+x[i+1])
		}

		sum = 0
		for i := 0; i < 1000000; i += 100 {
			sum += y[i]
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("Elapsed (sec) %s\n", elapsed)
	fmt.Println(sum)
}
