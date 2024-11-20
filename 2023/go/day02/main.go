package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const AocRootDir = "../.."
const Day = "day02"

var PuzzleInput = filepath.Join(AocRootDir, "inputs", Day, "input.txt")

func main() {
	puzzleFile, err := filepath.Abs(PuzzleInput)
	if err != nil {
		log.Fatalf("error reading puzzle input: %v", err)
	}
	fmt.Println("Puzlle input:", puzzleFile)
	data, _ := os.ReadFile(PuzzleInput)
	lines := strings.Split(string(data), "\n")

	fmt.Println("\n--- PART 1 ---")
	part1(lines)

	fmt.Println("\n--- PART 2 ---")
	part2(lines)
}

func part1(lines []string) {
	var possibleGames []int

	LIMITS := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for _, game := range lines {
		// ignore emtpy lines, e.g. last line
		if len(game) == 0 {
			continue
		}

		possible := true
		// fmt.Println(game)
		parts := strings.Split(game, ":")
		gameId, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])

		draws := strings.Split(parts[1], "; ")
		for _, draw := range draws {

			cubes := strings.Split(draw, ", ")

			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")
				// fmt.Println("   cube: " + cube)
				cubeParts := strings.Split(cube, " ")
				count, _ := strconv.Atoi(cubeParts[0])
				color := cubeParts[1]

				if count > LIMITS[color] {
					// fmt.Printf("      Game %d IMPOSSIBLE (%s count is %d)", gameId, color, count)
					possible = false
					break
				}
			}
		}
		if possible {
			// The below is used becuase Go lacks a set data structure
			// Check if the game id is already a possible game
			present := false
			for _, val := range possibleGames {
				if val == gameId {
					present = true
				}
			}
			// Only add if not present
			if !present {
				possibleGames = append(possibleGames, gameId)
			}
		}
	}
	// fmt.Println(possibleGames)
	sum := 0
	for _, id := range possibleGames {
		sum += id
	}
	fmt.Printf("Part 1 answer: %d\n", sum)
}

func part2(lines []string) {

	// {"gameId": {"color": count}, ...}
	counts := make(map[int]map[string]int)

	for _, game := range lines {
		// ignore emtpy lines, e.g. last line
		if len(game) == 0 {
			continue
		}

		parts := strings.Split(game, ":")
		gameId, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])

		// map for {"color": count, ...}
		if counts[gameId] == nil {
			counts[gameId] = make(map[string]int)
		}

		draws := strings.Split(parts[1], "; ")
		for _, draw := range draws {

			cubes := strings.Split(draw, ", ")

			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")
				cubeParts := strings.Split(cube, " ")
				count, _ := strconv.Atoi(cubeParts[0])
				color := cubeParts[1]

				if counts[gameId][color] < count {
					counts[gameId][color] = count
				}
			}
		}
	}
	// sum the powers
	sum := 0
	for _, gameCounts := range counts {
		product := 1 // NOTE: 1, not 0 (we're multiplying!)
		for _, val := range gameCounts {
			product *= val
		}
		sum += product
	}
	fmt.Println(sum)
}
