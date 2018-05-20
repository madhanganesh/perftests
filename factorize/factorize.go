package main

import (
  "fmt"
	"time"
)

func factors(value int64) []int64 {
  numbers := []int64{}
  var i int64 = 1
  for i=1; i<=value; i++ {
    if (value % i == 0) {
      numbers = append(numbers, i)
    }
  }
  return numbers
}

func main() {
  var input int64 = 100000000
	start := time.Now()
  data := factors(input)
  fmt.Println(len(data))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
