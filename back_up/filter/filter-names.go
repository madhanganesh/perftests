package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func filter(vs []string, f func(string) bool) []string {
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
	originalCount := len(names)

	start := time.Now()

	names = filter(names, func(name string) bool { return !strings.Contains(name, "A") })

	elapsed := time.Since(start)
	fmt.Println("original count is", originalCount)
	fmt.Println("filtered count is", len(names))
	fmt.Println(elapsed)
}
