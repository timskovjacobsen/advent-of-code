package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

const AocRootDir = "../.."
const Day = "day04"

var PuzzleInput = filepath.Join(AocRootDir, "inputs", Day, "input.txt")

func main() {
	puzzleFile, err := filepath.Abs(PuzzleInput)
	if err != nil {
		log.Fatalf("error reading puzzle input: %v", err)
	}
	fmt.Println("Puzlle input:", puzzleFile)
	data, _ := os.ReadFile(puzzleFile)
	// data, _ := os.ReadFile("test_input.txt")
	rawLines := strings.Split(string(data), "\n")

	lines := []string{}
	for _, line := range rawLines {
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	part1(lines)
	part2(lines)
}

func findWinners(line string) mapset.Set[int] {
	// fmt.Println(line)
	card := strings.Split(line, ": ")

	numbers := strings.Split(card[1], "|")
	winningNums := strings.Fields(numbers[0])

	myNums := strings.Fields(numbers[1])

	winningSet := mapset.NewSet[int]()
	mySet := mapset.NewSet[int]()

	for _, winningNum := range winningNums {
		winningNum, _ := strconv.Atoi(winningNum)
		winningSet.Add(winningNum)
	}
	for _, myNum := range myNums {
		myNum, _ := strconv.Atoi(myNum)
		mySet.Add(myNum)
	}
	myWinners := winningSet.Intersect(mySet)
	return myWinners
}

func part1(lines []string) {

	res := 0
	for _, line := range lines {
		myWinners := findWinners(line)
		// fmt.Println(myWinners)

		cardWorth := 0
		if myWinners.Cardinality() > 0 {
			cardWorth = 1
			for i := 0; i < myWinners.Cardinality()-1; i++ {
				cardWorth *= 2
			}
		}
		res += cardWorth
	}
	fmt.Println("PART 1: ", res)
}

func iterate(nWinnersMap map[int]int, cardCountsMap map[int]int) int {
	totalCards := 0
	cardId := 1
	for len(cardCountsMap) > 0 {
		nWinners := nWinnersMap[cardId]

		for cardCountsMap[cardId] > 0 {
			for i := cardId + 1; i <= cardId+nWinners; i++ {
				cardCountsMap[i] += 1
			}
			totalCards += 1
			cardCountsMap[cardId] -= 1
		}
		delete(cardCountsMap, cardId)
		cardId++
	}
	return totalCards
}

func recursiveHelper(lines []string, i int) int {
	// here we deal with all the newly spawned scratchcards and "scratch" them
	myWinners := findWinners(lines[i])
	spawnedTotal := 0

	// spawn cards in the following lines equal to amount of winning numbers
	for j := 0; j < myWinners.Cardinality(); j++ {
		nextLineIdx := i + 1
		spawnedTotal += recursiveHelper(lines, nextLineIdx+j)
	}
	// the card that called this function contributes one to the sum
	// we also add the total of its spawned cards
	return 1 + spawnedTotal
}

func recurse(lines []string) int {
	// recurse through all the lines
	// the helper will deal with the "scratching" cards that are generated as
	// the runout play out
	sum := 0
	for i := 0; i < len(lines); i++ {
		sum += recursiveHelper(lines, i)
	}
	return sum
}

func part2(lines []string) {
	nWinnersMap := make(map[int]int)

	// Map representing the current pool of scratchcards
	cardCountsMap := make(map[int]int)

	for _, line := range lines {
		//
		myWinners := findWinners(line)
		nWinners := myWinners.Cardinality()
		cardField := strings.Split(line, ":")[0]
		cardIdStr := strings.Fields(cardField)[1]
		cardId, _ := strconv.Atoi(cardIdStr)
		nWinnersMap[cardId] = nWinners

		// All cards have a count of one in the beginning
		cardCountsMap[cardId] = 1
	}

	fmt.Println("PART 2:")

	totalCardsIterate := iterate(nWinnersMap, cardCountsMap)
	fmt.Println("  Iterative: ", totalCardsIterate)

	totalCardsRecurse := 0
	totalCardsRecurse = recurse(lines)
	fmt.Println("  Recursive: ", totalCardsRecurse)
}
