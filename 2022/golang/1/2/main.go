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
	topThree := []int{0, 0, 0}
	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
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

		sum += num
	}

	log.Println(topThree[0] + topThree[1] + topThree[2])
}
