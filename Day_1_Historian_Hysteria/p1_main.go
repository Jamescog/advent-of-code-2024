package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
		}
		r, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting left number:", err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	totalOffset := 0
	for i := 0; i < len(left); i++ {
		offset := absInt(left[i] - right[i])
		totalOffset += offset
	}

	fmt.Println(totalOffset)

}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
