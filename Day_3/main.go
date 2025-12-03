package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RESPONSE PART 1 : 17613
// RESPONSE PART 2 : 175304218462560

func main() {

	testRanges, err := FileReader("C:\\Dev\\repo_git\\AOC2025\\Day_3\\input\\input.txt")
	if err != nil {
		fmt.Printf("ERROR convert File : %v", err)
	}

	// testRanges := []string{
	// 	"987654321111111",
	// 	"811111111111119",
	// 	"234234234234278",
	// 	"818181911112111",
	// }

	r := SumVoltages(testRanges)
	fmt.Printf("RESPONSE : %d", r)
}

func SumVoltages(banks []string) int {
	sum := 0
	for _, b := range banks {
		v := maxVoltage(b)
		sum += v
	}
	return sum
}

func maxVoltage(s string) int {
	// Variables
	max1, max2 := 0, 0

	for i, c := range s {

		num, _ := strconv.Atoi(string(c))

		if num > max1 && i != len(s)-1 {
			max1 = num
			max2 = 0
			continue
		}

		if num > max2 {
			max2 = num
		}
	}

	first := strconv.Itoa(max1)
	second := strconv.Itoa(max2)
	result, _ := strconv.Atoi(first + second)

	return result
}

func maxVoltagePart2(s string) int {
	// Variables
	var maxMap [12]int
	currentMax := 0

	for i, c := range s {

		num, _ := strconv.Atoi(string(c))
		for j, maxj := range maxMap {

			if num > maxj && len(s)-i+1 > (12-j) {
				maxMap[j] = num
				maxMap = ResetFrom(maxMap, j+1)
				break
			}

			if maxj == 0 {
				maxMap[j] = num
				currentMax++
				break
			}
		}
	}

	tempResult := ""
	for _, max := range maxMap {
		tempResult += strconv.Itoa(max)
	}
	result, _ := strconv.Atoi(tempResult)

	return result
}

func ResetFrom(maxMap [12]int, i int) [12]int {
	for j := i; j < len(maxMap); j++ {
		maxMap[j] = 0
	}
	return maxMap
}

func FileReader(path string) ([]string, error) {
	var result []string

	file, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" { // ignore les lignes vides
			result = append(result, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}
