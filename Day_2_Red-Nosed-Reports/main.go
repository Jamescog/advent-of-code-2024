package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func toIntArry(sa []string) []int {
	intArry := []int{}
	for _, num := range strings.Fields(sa[0]) {
		intVal, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("Error parsing string %v", err.Error())
		}
		intArry = append(intArry, intVal)
	}
	return intArry
}

func readReports(fname string) [][]int {
	reports := [][]int{}
	file, err := os.Open(fname)

	if err != nil {
		log.Fatalf("Error readining file %v: %v", fname, err.Error())
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Split(line, "	")
		reports = append(reports, toIntArry(values))
	}

	return reports
}

func isGradual(ia []int) bool {

	flow := []string{}

	for i := 0; i < len(ia)-1; i++ {
		t := ""
		if ia[i+1] < ia[i] {
			t = "dec"
		} else {
			t = "inc"
		}
		if !slices.Contains(flow, t) {
			flow = append(flow, t)
		}
		if len(flow) > 1 {
			return false
		}
	}

	return true
}

func absX(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func isValidDiff(ia []int) bool {
	for i := 0; i < len(ia)-1; i++ {
		diff := absX(ia[i+1] - ia[i])

		if !(diff >= 1 && diff <= 3) {
			return false
		}

	}
	return true
}

func main() {
	fname := "input.txt"
	var safeCount int
	ias := readReports(fname)

	for _, ia := range ias {

		if isGradual(ia) && isValidDiff(ia) {
			safeCount++
		}
	}

	fmt.Println("Safe count:", safeCount)
}
