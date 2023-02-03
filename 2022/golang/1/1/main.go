package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputString := string(input)

	lines := strings.Split(inputString, "\n")

	sum := 0
	biggest := 0

	for _, line := range lines {
		if line == "" {
			sum = 0
			continue
		}

		num, _ := strconv.Atoi(line)
		sum += num

		if sum > biggest {
			biggestGroup = sum
		}
	}

	log.Println(biggest)
}
