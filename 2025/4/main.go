package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, err := readGridFromFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading grid: %v\n", err)
		return
	}
	printInitialMovableCount(grid)
	printTotalMovable(grid)
}

func printInitialMovableCount(grid [][]string) {
	sum := 0

	for r, row := range grid {
		for c, ch := range row {
			if ch == "@" {
				count := getAdjacentCount(grid, r, c)
				if count < 4 {
					sum++
				}
			}
		}
	}
	fmt.Printf("Initially, %v are moveable\n", sum)
}

func printTotalMovable(grid [][]string) {
	moved := 0
	prevMoved := -1
	for moved != prevMoved {
		prevMoved = moved
		for r, row := range grid {
			for c, ch := range row {
				if ch == "@" {
					count := getAdjacentCount(grid, r, c)
					if count < 4 {
						moved++
						grid[r][c] = "."
					}
				}
			}
		}
	}
	fmt.Printf("In total, %v were moved\n", moved)
}

func getAdjacentCount(grid [][]string, row, col int) int {
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	for _, d := range dirs {
		r := row + d[0]
		c := col + d[1]

		if r < 0 || r >= rows || c < 0 || c >= cols {
			continue
		}

		if grid[r][c] == "@" {
			count++
		}
	}

	return count
}

func readGridFromFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		chars := make([]string, len(line))
		for i, r := range line {
			chars[i] = string(r)
		}
		grid = append(grid, chars)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error during scanning: %w", err)
	}

	return grid, nil
}
