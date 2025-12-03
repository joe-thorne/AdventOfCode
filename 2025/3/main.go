package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	sumA := 0
	sumB := 0

	for scanner.Scan() {
		line := scanner.Text()
		sumA += solveLine(line, 2)
		sumB += solveLine(line, 12)
	}
	fmt.Printf("Sum for Part 1 is %d\n", sumA)
	fmt.Printf("Sum for Part 2 is %d\n", sumB)
}

func solveLine(line string, d int) int {
	lineJolts := 0
	start := 0
	end := len(line) - (d - 1)
	for range d {
		max := 0
		for j := start; j < end; j++ {
			n, err := strconv.Atoi(string(line[j]))
			if err != nil {
				panic(err)
			}
			if n > max {
				max = n
				start = j + 1
			}
		}
		end++
		lineJolts = (10 * lineJolts) + max
	}
	return lineJolts
}
