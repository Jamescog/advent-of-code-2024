package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		l, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting left number:", err)
			continue
		}
		r, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting right number:", err)
			continue
		}

		left = append(left, l)
		right = append(right, r)
	}

	similarityScore := 0
	for _, n := range left {
		similarityScore += countOccurance(right, n)
	}
	fmt.Println(similarityScore)
}

func countOccurance(r []int, m int) int {
	count := 0
	for _, n := range r {
		if n == m {
			count++
		}
	}
	return m * count
}
