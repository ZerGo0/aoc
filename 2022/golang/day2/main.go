package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	var part int
	flag.IntVar(&part, "p", 1, "part 1 or 2")
	flag.Parse()

	if part != 1 && part != 2 {
		log.Fatal("invalid part")
	}

	if part == 1 {
		part1()
	} else {
		part2()
	}

}

func part1() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	// Rules:
	// Play     Points Enemy Me
	// Rock     1      A     X
	// Paper    2      B     Y
	// Scissors 3      C     Z
	//
	// Lose = 1 Point
	// Draw = 3 Points
	// Win  = 6 Points

	plays := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	playPoints := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	counterPlays := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	totalScore := 0

	for sc.Scan() {
		matchPlays := strings.Split(sc.Text(), " ")

		score := playPoints[matchPlays[1]]

		if matchPlays[1] == counterPlays[matchPlays[0]] {
			// Win
			totalScore += score + 6
		} else if matchPlays[1] == plays[matchPlays[0]] {
			// Draw
			totalScore += score + 3
		} else {
			// Lose
			totalScore += score
		}
	}

	log.Println(totalScore)
}

func part2() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	playPoints := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	outcomePoints := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	plays := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	losePlays := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	counterPlays := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	totalScore := 0

	for sc.Scan() {
		matchPlays := strings.Split(sc.Text(), " ")

		totalScore += outcomePoints[matchPlays[1]]

		switch matchPlays[1] {
		case "X":
			// Lose
			totalScore += playPoints[losePlays[matchPlays[0]]]
		case "Y":
			// Draw
			totalScore += playPoints[plays[matchPlays[0]]]
		case "Z":
			// Win
			totalScore += playPoints[counterPlays[matchPlays[0]]]
		}

	}

	log.Println(totalScore)
}
