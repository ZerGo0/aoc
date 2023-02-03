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
	topThree := []int{0, 0, 0}
	for _, line := range lines {
		if len(line) == 0 {
			if sum > topThree[2] {
				// sum is greater than the smallest of the top three
				// now we need to find the right place to insert it
				for i := 0; i < 3; i++ {
					if sum > topThree[i] {
						// We need to shift the slice
						for j := 2; j > i; j-- {
							topThree[j] = topThree[j-1]
						}
						topThree[i] = sum
						break
					}
				}
			}

			sum = 0
			continue
		}

		num, _ := strconv.Atoi(line)
		sum += num
	}

	log.Println(topThree[0] + topThree[1] + topThree[2])
}
