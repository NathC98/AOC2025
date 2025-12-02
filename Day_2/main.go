package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RESPONSE PART 1 : 64215794229
// RESPONSE PART 2 : 85513235135

func main() {

	contentFile, err := FileReader("C:\\Dev\\repo_git\\AOC2025\\Day_2\\input\\input.txt")
	if err != nil {
		fmt.Printf("ERROR convert File : %v", err)
	}
	testRanges := strings.Split(contentFile[0], ",")

	// testRanges := []string{
	// 	"11-22",
	// 	"95-115",
	// 	"998-1012",
	// 	"1188511880-1188511890",
	// 	"222220-222224",
	// 	"1698522-1698528",
	// 	"446443-446449",
	// 	"38593856-38593862",
	// 	"565653-565659",
	// 	"824824821-824824827",
	// 	"2121212118-2121212124",
	// }

	r := SumNotIDsPart2(testRanges)
	fmt.Printf("RESPONSE : %d", r)
}

func SumNotIDs(ranges []string) int {
	sum := 0
	for _, r := range ranges {
		rangeSplited := strings.Split(r, "-")
		floorRange, _ := strconv.Atoi(rangeSplited[0])
		ceilingRange, _ := strconv.Atoi(rangeSplited[1])
		for id := floorRange; id <= ceilingRange; id++ {
			stringId := strconv.Itoa(id)
			if IsNotAnID(stringId) {
				sum += id
			}
		}
	}
	return sum
}

func IsNotAnID(s string) bool {
	return len(s)%2 == 0 && s[:len(s)/2] == s[len(s)/2:]
}

func SumNotIDsPart2(ranges []string) int {
	sum := 0
	for _, r := range ranges {
		rangeSplited := strings.Split(r, "-")
		floorRange, _ := strconv.Atoi(rangeSplited[0])
		ceilingRange, _ := strconv.Atoi(rangeSplited[1])
		for id := floorRange; id <= ceilingRange; id++ {
			stringId := strconv.Itoa(id)
			if IsNotAnIDPart2(stringId) {
				sum += id
			}
		}
	}
	return sum
}

func IsNotAnIDPart2(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 && strings.Count(s, s[:i]) == len(s)/i {
			return true
		}
	}
	return false
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
