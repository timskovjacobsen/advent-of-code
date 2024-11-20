package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const AocRootDir = "../../.."
const Year = "2023"
const Day = "day05"

var PuzzleInput = filepath.Join(AocRootDir, Year, "inputs", Day, "input.txt")

func main() {
	puzzleFile, err := filepath.Abs(PuzzleInput)
	if err != nil {
		log.Fatalf("error reading puzzle input: %v", err)
	}
	fmt.Println("Puzlle input:", puzzleFile)
	data, _ := os.ReadFile(puzzleFile)
	// data, _ := os.ReadFile("input.txt")
	rawLines := strings.Split(string(data), "\n")

	lines := []string{}
	for _, line := range rawLines {
		if len(line) != 0 {
			lines = append(lines, line)
		}
	}

	part1(lines)
}

func part1(lines []string) {
	// ..
}
