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
	fileName := "../../data/names.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error in openeing the data file. Check README.md file")
		os.Exit(1)
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
