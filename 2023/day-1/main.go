package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var INPUT_FILE = "input.txt"

func main() {
	sum1, err := getResult(true)
	if err != nil {
		fmt.Println(err)
	}

	sum2, err := getResult(false)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Part I: %d\n", sum1)
	fmt.Printf("Part II: %d\n", sum2)
}

// Reads file line by line and returns sum of the numbers
// consisted of first and last digits
//
// Args:
// file - file pointer
// numbersOnly - if false, also spelled digits will be used
//
// Returns:
// int - sum of the numbers
func getResult(numbersOnly bool) (int, error) {
	file, err := readFile(INPUT_FILE)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		var line string
		line = scanner.Text()

		firstDigit := findDigitInString(line, 1, numbersOnly)
		lastDigit := findDigitInString(line, -1, numbersOnly)

		twoDigitNumber := firstDigit*10 + lastDigit
		sum += twoDigitNumber
	}

	return sum, nil
}

// Reads file and returns file pointer
func readFile(inputFile string) (*os.File, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	return file, nil
}

// Finds first or last digit in the string
//
// Args:
// line - string to search
// direction - direction to search
// (1 - from left to right, -1 - from right to left)
//
// Returns:
// int - found digit
func findDigitInString(line string, direction int, numbersOnly bool) int {
	numberDictionary := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	if direction == 1 {
		for i := 0; i < len(line); i++ {
			if charIsDigit(string(line[i])) {
				number, _ := strconv.Atoi(string(line[i]))
				return number
			}

			if !numbersOnly {
				for key, _ := range numberDictionary {
					if strings.Contains(line[:i+1], key) {
						return numberDictionary[key]
					}
				}
			}

		}
	}

	if direction == -1 {
		for i := len(line) - 1; i >= 0; i-- {
			if charIsDigit(string(line[i])) {
				number, _ := strconv.Atoi(string(line[i]))
				return number
			}

			if !numbersOnly {
				for key, _ := range numberDictionary {
					if strings.Contains(line[i:], key) {
						return numberDictionary[key]
					}
				}
			}

		}
	}

	return 0
}

// Checks if given character is a digit
func charIsDigit(char string) bool {
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}

	return false
}
