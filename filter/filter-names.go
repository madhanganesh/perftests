package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Filter(vs []string, f func(string) bool) []string {
	vsf := vs[:0]
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func main() {
	file, err := os.Open("../data/names.txt")
	if err != nil {
		panic("error in opening the file")
	}

	names := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	fmt.Println("count is", len(names))
	start := time.Now()

	names = Filter(names, func(name string) bool { return !strings.Contains(name, "A") })

	elapsed := time.Since(start)
	fmt.Println(elapsed)

	fmt.Println("count is", len(names))
}
