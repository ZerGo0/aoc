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

	populateCrates := true
	crateRows := make(map[int]string)

	for sc.Scan() {
		line := sc.Text()

		// Skip if the line is empty or the second character is a 1
		if len(line) == 0 || line[1] == 49 {
			populateCrates = false
			continue
		}

		if populateCrates {
			index := 0
			for i := 1; i < len(line); i += 4 {
				char := line[i]
				index++

				if char == 32 {
					// There is no crate here
					continue
				}

				crateRows[index] += string(char)
			}
		} else {
			split := strings.Split(line, " ")

			amount, _ := strconv.Atoi(split[1])
			source, _ := strconv.Atoi(split[3])
			target, _ := strconv.Atoi(split[5])

			movingCrates := crateRows[source][0:amount]
			crateRows[source] = crateRows[source][amount:]

			// We need to reverse the crates and put them in front of the target
			reversed := ""
			for i := len(movingCrates) - 1; i >= 0; i-- {
				reversed += string(movingCrates[i])
			}

			crateRows[target] = reversed + crateRows[target]
		}
	}

	topCrates := ""
	for i := 1; i <= len(crateRows); i++ {
		topCrates += string(crateRows[i][0])
	}

	// TODO: Check why this is not correct???
	//for _, v := range crateRows {
	//	log.Println(v)
	//}

	log.Println(topCrates)
}

func part2() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	populateCrates := true
	crateRows := make(map[int]string)

	for sc.Scan() {
		line := sc.Text()

		// Skip if the line is empty or the second character is a 1
		if len(line) == 0 || line[1] == 49 {
			populateCrates = false
			continue
		}

		if populateCrates {
			index := 0
			for i := 1; i < len(line); i += 4 {
				char := line[i]
				index++

				if char == 32 {
					// There is no crate here
					continue
				}

				crateRows[index] += string(char)
			}
		} else {
			split := strings.Split(line, " ")

			amount, _ := strconv.Atoi(split[1])
			source, _ := strconv.Atoi(split[3])
			target, _ := strconv.Atoi(split[5])

			movingCrates := crateRows[source][0:amount]
			crateRows[source] = crateRows[source][amount:]

			crateRows[target] = movingCrates + crateRows[target]
		}
	}

	topCrates := ""
	for i := 1; i <= len(crateRows); i++ {
		topCrates += string(crateRows[i][0])
	}

	// TODO: Check why this is not correct???
	//for _, v := range crateRows {
	//	log.Println(v)
	//}

	log.Println(topCrates)
}
