package main

import (
	"bufio"
	"flag"
	"log"
	"os"
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

	combined := 0
	for sc.Scan() {
		line := sc.Text()

		firstDigit := 0
		lastDigit := 0
		// find the first number in the current line
	first:
		for i := 0; i < len(line); i++ {
			x := line[i]

			if x >= '0' && x <= '9' {
				if firstDigit == 0 {
					firstDigit = int(x - '0')
					lastDigit = firstDigit
					break first
				}
			}
		}

	last:
		for i := len(line) - 1; i >= 0; i-- {
			x := line[i]

			if x >= '0' && x <= '9' {
				lastDigit = int(x - '0')
				break last
			}
		}

		combined += firstDigit*10 + lastDigit
	}

	log.Printf("combined: %d", combined)
}

func part2() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	numberTexts := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	combined := 0
	for sc.Scan() {
		line := sc.Text()

		firstDigit := 0
		lastDigit := 0

		// find the first number in the current line
	first:
		for i := 0; i < len(line); i++ {
			x := line[i]

			// Simple number check
			if x >= '0' && x <= '9' {
				if firstDigit == 0 {
					firstDigit = int(x - '0')
					lastDigit = firstDigit
					break first
				}
			}

			// The numbers can also be a word now
			// We can just check if the byte char is one of the numberTexts
			if x == 'o' || x == 't' || x == 'f' || x == 's' || x == 'e' || x == 'n' {
				for j := 0; j < len(numberTexts); j++ {
					word := numberTexts[j]
					if x == word[0] {
						// Check if the length of the word is within the bounds of the line
						if i+len(word) <= len(line) {
							// We need to match the entire word
							if line[i:i+len(word)] == word {
								if firstDigit == 0 {
									firstDigit = j + 1
									lastDigit = firstDigit
									break first
								}
							}
						}
					}
				}
			}
		}

	last:
		for i := len(line) - 1; i >= 0; i-- {
			x := line[i]

			// Simple number check
			if x >= '0' && x <= '9' {
				lastDigit = int(x - '0')
				break last
			}

			// The numbers can also be a word now
			// We can just check if the byte char is one of the numberTexts
			if x == 'o' || x == 't' || x == 'f' || x == 's' || x == 'e' || x == 'n' {
				for j := 0; j < len(numberTexts); j++ {
					word := numberTexts[j]

					if x == word[0] {
						// Check if the length of the word is within the bounds of the line
						if i+len(word) <= len(line) {
							// We need to match the entire word
							if line[i:i+len(word)] == word {
								lastDigit = j + 1
								break last
							}
						}
					}
				}
			}
		}

		combined += firstDigit*10 + lastDigit
	}

	log.Printf("combined: %d", combined)
}
