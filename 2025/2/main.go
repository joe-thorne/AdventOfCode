package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	endOfFileReached := false
	sumA := 0
	sumB := 0
	for !endOfFileReached {
		token, err := reader.ReadString(',')
		if err != nil {
			if err == io.EOF {
				endOfFileReached = true
			} else {
				fmt.Printf("Error reading from file: %v\n", err)
			}
		}
		sequence := strings.Split(token, "-")
		start, err := strconv.Atoi(sequence[0])
		if err != nil {
			panic(err)
		}
		endingString := sequence[1]
		if endingString[len(endingString)-1] == ',' {
			endingString = endingString[0:(len(endingString) - 1)]
		}
		end, err := strconv.Atoi(endingString)
		if err != nil {
			panic(err)
		}
		for i := start; i <= end; i++ {
			if isInvalidA(i) {
				sumA += i
			}
			if isInvalidB(i) {
				sumB += i
			}

		}
	}
	fmt.Printf("The sum of Part 1 is %d\n", sumA)
	fmt.Printf("The sum of Part 2 is %d\n", sumB)
}

func isInvalidA(num int) bool {
	s := fmt.Sprint(num)
	length := len(s)
	if length%2 == 0 {
		if s[0:length/2] == s[length/2:] {
			return true
		}
	}
	return false
}

func isInvalidB(num int) bool {
	s := fmt.Sprint(num)
	length := len(s)
	factors := getFactors(length)
	for _, fac := range factors {
		isInvalid := true
		for i := fac; i < length; i += fac {
			if s[i-fac:i] != s[i:i+fac] {
				isInvalid = false
				break
			}
		}
		if isInvalid {
			return true
		}
	}
	return false
}

func getFactors(n int) []int {
	if n <= 1 {
		return []int{}
	}
	factors := []int{}
	limit := int(math.Sqrt(float64(n)))
	factors = append(factors, 1)
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if i*i != n {
				factors = append(factors, n/i)
			}
		}
	}

	return factors
}
