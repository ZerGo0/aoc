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

	regX := 1
	cycle := 0
	cycleValues := make([]int, 6)
	for sc.Scan() {
		line := sc.Text()

		split := strings.Split(line, " ")

		addCycles := 0
		switch split[0] {
		case "noop":
			addCycles = 1
		case "addx":
			addCycles = 2
		}

		for i := 0; i < addCycles; i++ {
			cycle++
			if cycle == 20 ||
				cycle == 60 ||
				cycle == 100 ||
				cycle == 140 ||
				cycle == 180 ||
				cycle == 220 {
				log.Println(cycle, regX, cycle*regX)
				cycleValues = append(cycleValues, cycle*regX)
			}
		}

		if split[0] == "addx" {
			value, _ := strconv.Atoi(split[1])
			regX += value
		}
	}

	sum := 0
	for _, v := range cycleValues {
		sum += v
	}

	log.Println(regX)
	log.Println(sum)
}

func part2() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	crtWidth := 40
	lines := make([]string, 0)
	lines = append(lines, "")

	regX := 1
	cycle := 0
	//p:
	for sc.Scan() {
		line := sc.Text()

		split := strings.Split(line, " ")

		addCycles := 0
		switch split[0] {
		case "noop":
			addCycles = 1
		case "addx":
			addCycles = 2
		}

		for i := 0; i < addCycles; i++ {
			cycle++

			currentRow := len(lines) - 1
			if len(lines[currentRow]) == crtWidth {
				currentRow++
				lines = append(lines, "")
				if len(lines) > 3 {
					//break p
				}
			}

			currentPos := cycle%crtWidth - 1  // We start at 0
			spritePosition := regX % crtWidth // We start at 1 because this is the middle pixel...
			delta := currentPos - spritePosition

			if delta >= -1 && delta <= 1 {
				lines[currentRow] += "#"
			} else {
				lines[currentRow] += "."
			}
		}

		if split[0] == "addx" {
			value, _ := strconv.Atoi(split[1])
			regX += value
		}
	}

	for _, l := range lines {
		log.Println(l)
	}
}
