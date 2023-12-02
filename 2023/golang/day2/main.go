package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
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

	possibleGamesSum := 0
	for sc.Scan() {
		text := sc.Text()
		textSplit := strings.Split(text, " ")

		gameId, err := strconv.Atoi(textSplit[1][0 : len(textSplit[1])-1])
		if err != nil {
			log.Fatal(err)
		}

		cubes := make([]int, 3)
		currentCubeCount := 0
	gameEval:
		for i := 2; i < len(textSplit); i++ {
			cube := textSplit[i]
			// if the number is even, it's a number of cubes
			// if the number is odd, it's the number of cubes and a comma or a semicolon
			if i%2 == 0 {
				cubeCount, err := strconv.Atoi(cube)
				if err != nil {
					log.Fatal(err)
				}

				currentCubeCount = cubeCount
			} else {
				color := cube[0]

				var colorIndex int
				var colorLimit int
				switch color {
				case 'r':
					colorIndex = 0
					colorLimit = 12
				case 'g':
					colorIndex = 1
					colorLimit = 13
				case 'b':
					colorIndex = 2
					colorLimit = 14
				}

				if currentCubeCount > colorLimit {
					break gameEval
				}

				cubes[colorIndex] += currentCubeCount

				if i == len(textSplit)-1 {
					possibleGamesSum += gameId
				}
			}
		}
	}

	log.Printf("sum of possible games: %d", possibleGamesSum)
}

func part2() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	gameResults := 0
	for sc.Scan() {
		text := sc.Text()
		textSplit := strings.Split(text, " ")

		cubes := make([]int, 3)
		currentCubeCount := 0
		for i := 2; i < len(textSplit); i++ {
			cube := textSplit[i]
			// if the number is even, it's a number of cubes
			// if the number is odd, it's the number of cubes and a comma or a semicolon
			if i%2 == 0 {
				cubeCount, err := strconv.Atoi(cube)
				if err != nil {
					log.Fatal(err)
				}

				currentCubeCount = cubeCount
			} else {
				color := cube[0]

				var colorIndex int
				switch color {
				case 'r':
					colorIndex = 0
				case 'g':
					colorIndex = 1
				case 'b':
					colorIndex = 2
				}

				if currentCubeCount > cubes[colorIndex] {
					cubes[colorIndex] = currentCubeCount
				}

				if i == len(textSplit)-1 {
					gameResults += cubes[0] * cubes[1] * cubes[2]
				}
			}
		}
	}

	log.Printf("game results: %d", gameResults)
}
