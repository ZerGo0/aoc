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

	total := 0

	for sc.Scan() {
		line := sc.Text()

		items := [2]string{
			line[:len(line)/2],
			line[len(line)/2:],
		}

		// TODO: Is there a better way to do this?
		// The int feels kind of useless, but the key lookup is faster than
		// iterating over a slice
		seen := initSeen(2)

		// TODO: ^
		commonItems := make(map[byte]bool)
		for i := 0; i < len(items[0]); i++ {
			seen[0][items[0][i]] = true
			seen[1][items[1][i]] = true

			if seen[0][items[1][i]] {
				if !commonItems[items[1][i]] {
					commonItems[items[1][i]] = true
				}
			}

			if seen[1][items[0][i]] {
				if !commonItems[items[0][i]] {
					commonItems[items[0][i]] = true
				}
			}
		}

		for char := range commonItems {
			if char >= 65 && char <= 90 {
				total += int(char) - 64 + 26
			} else {
				total += int(char) - 96
			}
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

	groupItems := [3]string{}
	index := 0

	for sc.Scan() {
		line := sc.Text()

		groupItems[index] = line
		index++

		if index != 3 {
			continue
		}

		seen := initSeen(3)
		commonItems := make(map[byte]bool)

		itemLengths := make([]int, len(groupItems))
		longestItem := lItem{}
		for i := 0; i < len(groupItems); i++ {
			itemLengths[i] = len(groupItems[i])
			if len(groupItems[i]) > longestItem.length {
				longestItem = lItem{
					index:  i,
					length: len(groupItems[i]),
				}
			}
		}

		for i := 0; i < longestItem.length; i++ {
			for j := 0; j < len(groupItems); j++ {
				if i < itemLengths[j] {
					seen[j][groupItems[j][i]] = true

					// If we've already seen this item, skip it
					if commonItems[groupItems[j][i]] {
						continue
					}

					// Check if we've seen this item in the other items
					seenInOthers := []int{}
					for k := 0; k < len(groupItems); k++ {
						if k != j && seen[k][groupItems[j][i]] {
							seenInOthers = append(seenInOthers, k)
						}
					}

					// If we've seen this item in all other items, add it to the common items
					if len(seenInOthers) == 2 {
						if !commonItems[groupItems[j][i]] {
							commonItems[groupItems[j][i]] = true
						}
					}
				}
			}
		}

		for char := range commonItems {
			if char >= 65 && char <= 90 {
				total += int(char) - 64 + 26
			} else {
				total += int(char) - 96
			}
		}

		groupItems = [3]string{}
		index = 0
	}

	log.Println(total)
}

type lItem struct {
	index  int
	length int
}

func initSeen(amount int) [](map[byte]bool) {
	seen := make([](map[byte]bool), amount)
	for i := 0; i < amount; i++ {
		seen[i] = make(map[byte]bool)
	}

	return seen
}
