package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type numRange struct {
	start int
	end   int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ranges := make([]numRange, 0)
	foundBlank := false
	sum := 0
	for scanner.Scan() {
		if !foundBlank {
			if scanner.Text() == "" {
				foundBlank = true
			} else {
				line := strings.Split(scanner.Text(), "-")
				start, err := strconv.Atoi(line[0])
				if err != nil {
					panic(err)
				}
				end, err := strconv.Atoi(line[1])
				if err != nil {
					panic(err)
				}
				ranges = append(ranges, numRange{start, end})
			}
		} else {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			for _, r := range ranges {
				if num >= r.start && num <= r.end {
					sum++
					break
				}
			}
		}
	}

	// Part 2
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})
	cMax := 0
	numValid := 0
	for _, r := range ranges {
		if r.end >= cMax {
			numValid += r.end - int(math.Max(float64(r.start), float64(cMax))) + 1
			cMax = r.end + 1
		}
	}

	fmt.Printf("%d good items\n", sum)
	fmt.Printf("%d valid IDs possible\n", numValid)

}
