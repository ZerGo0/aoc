package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	for sc.Scan() {
	}
}
