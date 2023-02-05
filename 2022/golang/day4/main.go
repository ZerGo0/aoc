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

	total := 0

	for sc.Scan() {
		line := sc.Text()

		assignments := strings.Split(line, ",")

		assignmentA := strings.Split(assignments[0], "-")
		assignmentB := strings.Split(assignments[1], "-")

		startA, _ := strconv.Atoi(assignmentA[0])
		endA, _ := strconv.Atoi(assignmentA[1])

		startB, _ := strconv.Atoi(assignmentB[0])
		endB, _ := strconv.Atoi(assignmentB[1])

		if startA >= startB && endA <= endB ||
			startB >= startA && endB <= endA {
			total++
		}
	}

	log.Println(total)
}

func part2() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	total := 0

	for sc.Scan() {
		line := sc.Text()

		assignments := strings.Split(line, ",")

		assignmentA := strings.Split(assignments[0], "-")
		assignmentB := strings.Split(assignments[1], "-")

		startA, _ := strconv.Atoi(assignmentA[0])
		endA, _ := strconv.Atoi(assignmentA[1])

		startB, _ := strconv.Atoi(assignmentB[0])
		endB, _ := strconv.Atoi(assignmentB[1])

		if startA <= startB && endA >= startB ||
			startB <= startA && endB >= startA {
			total++
		}
	}

	log.Println(total)
}
