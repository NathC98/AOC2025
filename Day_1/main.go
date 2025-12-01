package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RESPONSE PART 1 : 1064
// RESPONSE PART 2 : 6122

func main() {

	testDecoy, err := FileReader("C:\\Dev\\repo_git\\AOC2025\\Day_1\\input\\input.txt")
	if err != nil {
		fmt.Printf("ERROR convert File : %v", err)
	}

	// testDecoy := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}

	r, err := dialDecoyPart2(testDecoy)
	if err != nil {
		fmt.Printf("ERROR : %v", err)
	} else {
		fmt.Printf("RESPONSE : %d", r)

	}
}

func dialDecoyPart1(input []string) (int, error) {
	// Variables
	currentNumber := 50
	password := 0

	for _, move := range input {
		moveDirection := 1
		if move[0] == 'L' {
			moveDirection = -1
		}

		number := move[1:]
		moveNumber, err := strconv.Atoi(number)
		if err != nil {
			return 0, err
		}

		currentNumber = currentNumber + moveDirection*(moveNumber%100)
		if currentNumber < 0 {
			currentNumber = 100 + currentNumber
		}

		if currentNumber > 99 {
			currentNumber = currentNumber - 100
		}

		if currentNumber == 0 {
			password++
		}

	}

	return password, nil
}

func dialDecoyPart2(input []string) (int, error) {
	// Variables
	currentNumber := 50
	password := 0

	for _, move := range input {
		moveDirection := 1
		if move[0] == 'L' {
			moveDirection = -1
		}

		number := move[1:]
		moveNumber, err := strconv.Atoi(number)
		if err != nil {
			return 0, err
		}

		startFromZero := currentNumber == 0
		password = password + moveNumber/100
		currentNumber = currentNumber + moveDirection*(moveNumber%100)
		if currentNumber == 0 && !startFromZero {
			password++
		}

		if currentNumber < 0 {
			if !startFromZero {
				password++
			}
			currentNumber = 100 + currentNumber
		}

		if currentNumber > 99 {
			if !startFromZero {
				password++
			}
			currentNumber = currentNumber - 100
		}
	}

	return password, nil
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
