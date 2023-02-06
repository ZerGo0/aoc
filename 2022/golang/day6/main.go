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

	// There's only one line in the input
	sc.Scan()
	log.Println(findUniqueSeqIndex(sc.Text(), 4))
}

func part2() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	// There's only one line in the input
	sc.Scan()
	log.Println(findUniqueSeqIndex(sc.Text(), 14))
}

func findUniqueSeqIndex(text string, amount int) int {
	chars := make([]byte, 0, amount)
	for i := 0; i < len(text); i++ {
		// TODO: Is this faster than constantly looping over the next 4/14 chars?
		if found := arrayContains(chars, text[i]); found != -1 {
			chars = chars[found+1:]
			chars = append(chars, text[i])
			continue
		} else {
			if len(chars)+1 == amount {
				return i + 1
			}

			chars = append(chars, text[i])
		}
	}

	return -1
}

func arrayContains(arr []byte, b byte) int {
	for i, v := range arr {
		if v == b {
			return i
		}
	}

	return -1
}
