package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const AocRootDir = "../../.."
const Year = "2023"
const Day = "day01"

var PuzzleInput = filepath.Join(AocRootDir, Year, "inputs", Day, "input.txt")

func main() {
	puzzleFile, err := filepath.Abs(PuzzleInput)
	if err != nil {
		log.Fatalf("error reading puzzle input: %v", err)
	}
	fmt.Println("Puzlle input:", puzzleFile)
	part1()
	part2()
}

func part1() {
	data, _ := os.ReadFile(PuzzleInput)
	lines := strings.Split(string(data), "\n")

	sum := 0
	for _, line := range lines {

		digits := []string{}
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, string(char))
			}
		}
		if len(digits) > 0 {
			calibrationNum := fmt.Sprintf("%s%s", digits[0], digits[len(digits)-1])
			num, _ := strconv.Atoi(calibrationNum)
			sum += num
		}
	}

	fmt.Println(sum)
}

// getKeys returns the keys of a map as a slice of strings
func getKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// reverseKeys returns a map with reversed keys
func reverseKeys(m map[string]string) map[string]string {
	reversed := make(map[string]string, len(m))
	for k, v := range m {
		reversed[reverseString(k)] = v
	}
	return reversed
}

// reverseString reverses a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func part2() {
	m := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	// data, _ := os.ReadFile("test_input_part2.txt")
	data, _ := os.ReadFile(PuzzleInput)
	lines := strings.Split(string(data), "\n")

	// Build a pattern that matches any of the spelled-out digits
	pattern := "(\\d|" + strings.Join(getKeys(m), "|") + ")"
	reversedMap := reverseKeys(m)
	reversePattern := "(\\d|" + strings.Join(getKeys(reversedMap), "|") + ")"

	re := regexp.MustCompile(pattern)
	reReverse := regexp.MustCompile(reversePattern)

	sum := 0
	for _, line := range lines {

		firstIdx := re.FindStringIndex(line)
		lastIdx := reReverse.FindStringIndex(reverseString(line))
		first := ""
		last := ""
		if len(firstIdx) != 0 {
			first = line[firstIdx[0]:firstIdx[1]]
		}
		if len(lastIdx) != 0 {
			last = reverseString(line)[lastIdx[0]:lastIdx[1]]
		}
		var firstNum, lastNum string
		if firstVal, exists := m[first]; exists {
			firstNum = firstVal
		} else {
			firstNum = first
		}

		if lastVal, exists := reversedMap[last]; exists {
			lastNum = lastVal
		} else {
			lastNum = last
		}
		num, _ := strconv.Atoi(firstNum + lastNum)
		sum += num
	}

	fmt.Println(sum)
}
