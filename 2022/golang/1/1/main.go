package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	sum := 0
	biggest := 0

	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			sum = 0
			continue
		}

		sum += num

		if sum > biggest {
			biggest = sum
		}
	}

	log.Println(biggest)
}

