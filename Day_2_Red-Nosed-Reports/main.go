package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func toIntArray(sa []string) []int {
	intArr := make([]int, len(sa))
	for i, num := range sa {
		intVal, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("Error parsing string %v", err.Error())
		}
		intArr[i] = intVal
	}
	return intArr
}

func readReports(fname string) [][]int {
	reports := [][]int{}
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error reading file %v: %v", fname, err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		reports = append(reports, toIntArray(values))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}
	return reports
}

func isGradual(ia []int) bool {
	if len(ia) < 2 {
		return true
	}

	isIncreasing := ia[1] > ia[0]
	for i := 1; i < len(ia); i++ {
		if (ia[i] > ia[i-1]) != isIncreasing {
			return false
		}
	}
	return true
}

func absX(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func isValidDiff(ia []int) bool {
	for i := 0; i < len(ia)-1; i++ {
		diff := absX(ia[i+1] - ia[i])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isSafeWithDampner(ia []int) bool {
	for i := 0; i < len(ia); i++ {
		candidate := append([]int{}, ia[:i]...)
		candidate = append(candidate, ia[i+1:]...)

		if isGradual(candidate) && isValidDiff(candidate) {
			return true
		}
	}
	return false
}

func main() {
	fname := "input.txt"
	var safeCount, safeCountWithDampner int
	ias := readReports(fname)

	for _, ia := range ias {
		if isGradual(ia) && isValidDiff(ia) {
			safeCount++
		} else if isSafeWithDampner(ia) {
			safeCountWithDampner++
		}
	}

	fmt.Println("Safe count:", safeCount)
	fmt.Println("Safe count with dampner:", safeCount+safeCountWithDampner)
}
