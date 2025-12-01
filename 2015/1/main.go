package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer f.Close()

	scanner := bufio.NewReader(f)
	floor := 0
	charCount := 0
	var enteredBasement int
	hasBeenInBasement := false
	for {
		c, _, err := scanner.ReadRune()
		if err != nil {
			break
		}
		charCount++
		switch c {
		case '(':
			floor++
		case ')':
			floor--
			if floor < 0 && !hasBeenInBasement {
				enteredBasement = charCount
				hasBeenInBasement = true
			}
		}
	}
	fmt.Printf("Santa ended on Floor: %d\n", floor)
	fmt.Printf("Santa entered the basement on char: %d\n", enteredBasement)
}
