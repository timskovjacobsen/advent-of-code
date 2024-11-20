package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

const AocRootDir = "../../.."
const Year = "2023"
const Day = "day03"

var PuzzleInput = filepath.Join(AocRootDir, Year, "inputs", Day, "input.txt")

func main() {
	puzzleFile, err := filepath.Abs(PuzzleInput)
	if err != nil {
		log.Fatalf("error reading puzzle input: %v", err)
	}
	fmt.Println("Puzlle input:", puzzleFile)
	data, _ := os.ReadFile(puzzleFile)
	rawLines := strings.Split(string(data), "\n")

	// purge zero length lines
	lines := []string{}
	for _, line := range rawLines {
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	part1(lines)

	part2(lines)
}

// Return the adjacent eight locations of a point in 2D space
// Note that returned coordinates might be negative
func adjacentEight(x int, y int) [8][2]int {
	/* Visual example with P as the input (x, y) point and adjacent locations
	   numbered	(y-axis increases downwards)

			.........
			...123...
			...4P5...
			...678...
			.........  --> x
	*/
	var adjacent [8][2]int
	adjacent[0] = [2]int{x - 1, y - 1} // loc 1
	adjacent[1] = [2]int{x, y - 1}     // loc 2
	adjacent[2] = [2]int{x + 1, y - 1} // loc 3
	adjacent[3] = [2]int{x - 1, y}     // loc 4
	adjacent[4] = [2]int{x + 1, y}     // loc 5
	adjacent[5] = [2]int{x - 1, y + 1} // loc 6
	adjacent[6] = [2]int{x, y + 1}     // loc 7
	adjacent[7] = [2]int{x + 1, y + 1} // loc 8
	return adjacent
}

type Number struct {
	Value    string
	StartIdx int
	EndIdx   int
}

func extractNumbers(line string) []Number {
	var numbers []Number
	var currentNumber string
	var startIdx int

	for i, char := range line {
		if unicode.IsDigit(char) {
			// We have a digit, add to current number
			if currentNumber == "" {
				startIdx = i
			}
			currentNumber += string(char)
		} else {
			// Not a digit -> end current number if present
			if currentNumber != "" {
				numbers = append(numbers, Number{Value: currentNumber, StartIdx: startIdx, EndIdx: i - 1})
				currentNumber = ""
			}
		}
	}
	//
	if currentNumber != "" {
		numbers = append(numbers, Number{Value: currentNumber, StartIdx: startIdx, EndIdx: len(line) - 1})
	}
	return numbers
}

func part1(lines []string) {
	partNumbersSum := 0

	// (x, y) represent the point of interest that we are scanning around
	for y, line := range lines {
		// get the numbers in the line and the index of each digit
		numbers := extractNumbers(line)

		// if the number has multiple digits, we must scan around all of them
		// to properly detect an adjacent symbol
		for _, number := range numbers {
			// be sure not to double count numbers if multiple of its digits are
			// surrounded by a given symbol or multiple symbols
			alreadyAdded := false
			for x := number.StartIdx; x <= number.EndIdx; x++ {
				if alreadyAdded {
					break
				}
				// Scan all surrounding 8 locations, skip if they go out of board bounds
				for _, coords := range adjacentEight(x, y) {
					xx, yy := coords[0], coords[1]
					areCoordsPositive := xx >= 0 && yy >= 0
					areCoordsBeyondLen := xx < len(line) && yy < len(lines)
					if (areCoordsPositive) && (areCoordsBeyondLen) {
						char := lines[yy][xx]
						if !unicode.IsDigit(rune(char)) && char != '.' {
							// fmt.Printf("part number: %s\n", number.Value)
							numInt, _ := strconv.Atoi(number.Value)
							partNumbersSum += numInt
							alreadyAdded = true
						}
					}
				}
			}
		}
	}
	fmt.Println("PART 1: part number sum: ", partNumbersSum)
}

type Star struct {
	x           int
	y           int
	partNumbers []string
}

func part2(lines []string) {
	// Holds coordinates of * symbols as keys and its part numbers as values
	stars := map[[2]int][]string{}

	// (x, y) represent the point of interest that we are scanning around
	for y, line := range lines {
		// get the numbers in the line and the index of each digit
		numbers := extractNumbers(line)

		// if the number has multiple digits, we must scan around all of them
		// to properly detect an adjacent symbol
		for _, number := range numbers {
			// be sure not to double count numbers if multiple of its digits are
			// surrounded by a given symbol or multiple symbols
			alreadyAdded := false
			for x := number.StartIdx; x <= number.EndIdx; x++ {
				if alreadyAdded {
					break
				}
				// Scan all surrounding 8 locations, skip if they go out of board bounds
				for _, coords := range adjacentEight(x, y) {
					xx, yy := coords[0], coords[1]
					areCoordsPositive := xx >= 0 && yy >= 0
					areCoordsBeyondLen := xx < len(line) && yy < len(lines)
					if (areCoordsPositive) && (areCoordsBeyondLen) {
						char := lines[yy][xx]
						if char == '*' {
							// Add this number to the part numbers of this
							// * symbol
							_, ok := stars[coords]
							if ok {
								stars[coords] = append(stars[coords], number.Value)
							} else {
								stars[coords] = []string{number.Value}
							}
							alreadyAdded = true
						}
					}
				}
			}
		}
	}
	// The * symbols that have exactly two part numbers are gears
	// -> gets multiplied and added
	gearRatioSum := 0
	for _, partNums := range stars {
		if len(partNums) == 2 {
			p1, _ := strconv.Atoi(partNums[0])
			p2, _ := strconv.Atoi(partNums[1])
			gearRatio := p1 * p2
			gearRatioSum += gearRatio
		}
	}
	fmt.Println("PART 2: gear ratio sum: ", gearRatioSum)
}
