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

	value := 50
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("There was an error parsing %v: %v\n", line, err)
			os.Exit(1)
		}
		for i := 0; i < distance; i++ {
			switch direction {
			case "L":
				value--
			case "R":
				value++
			default:
				fmt.Printf("error reading direction \"%v\" from line %v", direction, line)
				os.Exit(1)
			}
			if value%100 == 0 {
				count++
			}
		}
	}
	fmt.Printf("Number of Zeros:  %d\n", count)

}
