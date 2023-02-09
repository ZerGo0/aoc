package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
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

	currentPos := make([][2]int, 2)
	currentPos[0][0] = 0
	currentPos[0][1] = 0

	currentPos[1][0] = 0
	currentPos[1][1] = 0

	// x and y position of the visited position with [x, y] as key and bool as value
	visitedPos := make(map[[2]int]bool)

	for sc.Scan() {
		line := sc.Text()

		direction := line[0]
		distance, _ := strconv.Atoi(line[2:])

		moveHead(currentPos, visitedPos, direction, distance)
	}

	//printGrid(visitedPos, currentPos)

	log.Println(len(visitedPos))
}

func part2() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	currentPos := make([][2]int, 10)
	currentPos[0][0] = 0
	currentPos[0][1] = 0

	// 1-9
	for i := 1; i < 10; i++ {
		currentPos[i][0] = 0
		currentPos[i][1] = 0
	}

	// x and y position of the visited position with [x, y] as key and bool as value
	// We only keep track of the last element of currentPos
	visitedPos := make(map[[2]int]bool)

	for sc.Scan() {
		line := sc.Text()

		direction := line[0]
		distance, _ := strconv.Atoi(line[2:])

		moveHead(currentPos, visitedPos, direction, distance)
	}

	//printGrid(visitedPos, currentPos)

	log.Println(len(visitedPos))
}

func moveHead(currentPos [][2]int, visitedPos map[[2]int]bool, dir byte, distance int) {
	for i := 0; i < distance; i++ {
		switch dir {
		// R
		case 82:
			currentPos[0][0]++
		// L
		case 76:
			currentPos[0][0]--
		// U
		case 85:
			currentPos[0][1]--
		// D
		case 68:
			currentPos[0][1]++
		}

		moveTail(currentPos, visitedPos, dir, distance)

		if _, ok := visitedPos[[2]int{currentPos[len(currentPos)-1][0], currentPos[len(currentPos)-1][1]}]; !ok {
			visitedPos[[2]int{currentPos[len(currentPos)-1][0], currentPos[len(currentPos)-1][1]}] = true
		}
	}
}

func moveTail(currentPos [][2]int, visitedPos map[[2]int]bool, dir byte, distance int) {
	for i := 1; i <= len(currentPos)-1; i++ {
		if currentPos[i-1][0]-currentPos[i][0] == 2 ||
			currentPos[i-1][0]-currentPos[i][0] == -2 ||
			currentPos[i-1][1]-currentPos[i][1] == 2 ||
			currentPos[i-1][1]-currentPos[i][1] == -2 {

			// Move up and right
			if currentPos[i-1][0] > currentPos[i][0] &&
				currentPos[i-1][1] < currentPos[i][1] {
				currentPos[i][0]++
				currentPos[i][1]--
				continue
			}

			// Move up and left
			if currentPos[i-1][0] < currentPos[i][0] &&
				currentPos[i-1][1] < currentPos[i][1] {
				currentPos[i][0]--
				currentPos[i][1]--
				continue
			}

			// Move down and right
			if currentPos[i-1][0] > currentPos[i][0] &&
				currentPos[i-1][1] > currentPos[i][1] {
				currentPos[i][0]++
				currentPos[i][1]++
				continue
			}

			// Move down and left
			if currentPos[i-1][0] < currentPos[i][0] &&
				currentPos[i-1][1] > currentPos[i][1] {
				currentPos[i][0]--
				currentPos[i][1]++
				continue
			}

			// Move up
			if currentPos[i-1][1] < currentPos[i][1] {
				currentPos[i][1]--
				continue
			}

			// Move down
			if currentPos[i-1][1] > currentPos[i][1] {
				currentPos[i][1]++
				continue
			}

			// Move right
			if currentPos[i-1][0] > currentPos[i][0] {
				currentPos[i][0]++
				continue
			}

			// Move left
			if currentPos[i-1][0] < currentPos[i][0] {
				currentPos[i][0]--
				continue
			}
		}
	}
}

func printGrid(visitedPos map[[2]int]bool, currentPos [][2]int) {
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for pos := range visitedPos {
		if pos[0] < minX {
			minX = pos[0]
		}

		if pos[0] > maxX {
			maxX = pos[0]
		}

		if pos[1] < minY {
			minY = pos[1]
		}

		if pos[1] > maxY {
			maxY = pos[1]
		}

	}

	// Print the grid
	for y := minY; y <= maxY; y++ {
		currentLine := ""
	xLoop:
		for x := minX; x <= maxX; x++ {
			for i := 1; i < len(currentPos); i++ {
				if currentPos[i][0] == x && currentPos[i][1] == y {
					currentLine += strconv.Itoa(i)
					continue xLoop
				}
			}

			if _, ok := visitedPos[[2]int{x, y}]; ok {
				currentLine += "#"
			} else {
				currentLine += "."
			}
		}

		log.Println(currentLine)
	}
}
