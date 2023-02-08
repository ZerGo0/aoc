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

	treeRows := make([]string, 0)
	for sc.Scan() {
		line := sc.Text()

		treeRows = append(treeRows, line)
	}

	visibleTrees := 0
	for i := 1; i < len(treeRows)-1; i++ {
		// Iterate over all characters in the row
		for j := 1; j < len(treeRows[i])-1; j++ {
			currentHeight := treeRows[i][j]

			// TODO: This could be optimized a lot...
			// We could keep track of the tallest tree in each direction
			// and only check if the current tree is taller than that
			// This way we could skip a lot of trees
			visibleSides := 4
			// Check left
			for k := j - 1; k >= 0; k-- {
				if treeRows[i][k] >= currentHeight {
					visibleSides--
					break
				}
			}

			// Check right
			for k := j + 1; k < len(treeRows[i]); k++ {
				if treeRows[i][k] >= currentHeight {
					visibleSides--
					break
				}
			}

			// Check up
			for k := i - 1; k >= 0; k-- {
				if treeRows[k][j] >= currentHeight {
					visibleSides--
					break
				}
			}

			// Check down
			for k := i + 1; k < len(treeRows); k++ {
				if treeRows[k][j] >= currentHeight {
					visibleSides--
					break
				}
			}

			if visibleSides != 0 {
				visibleTrees++
			}
		}
	}

	log.Println(visibleTrees + len(treeRows[0])*2 + len(treeRows)*2 - 4)
}

func part2() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	treeRows := make([]string, 0)
	for sc.Scan() {
		line := sc.Text()

		treeRows = append(treeRows, line)
	}

	tallestTree := 0
	for i := 1; i < len(treeRows)-1; i++ {
		// Iterate over all characters in the row
		for j := 1; j < len(treeRows[i])-1; j++ {
			currentHeight := treeRows[i][j]

			// We want to find out the distance the nearest tree that it the same height or higher
			// in each direction. We multiply it with totalScore to get the total score for this tree
			totalScore := 1
			// Check left
			for k := j - 1; k >= 0; k-- {
				if treeRows[i][k] >= currentHeight || k == 0 {
					totalScore *= j - k
					break
				}
			}

			// Check right
			for k := j + 1; k < len(treeRows[i]); k++ {
				if treeRows[i][k] >= currentHeight || k == len(treeRows[i])-1 {
					totalScore *= k - j
					break
				}
			}

			// Check up
			for k := i - 1; k >= 0; k-- {
				if treeRows[k][j] >= currentHeight || k == 0 {
					totalScore *= i - k
					break
				}
			}

			// Check down
			for k := i + 1; k < len(treeRows); k++ {
				if treeRows[k][j] >= currentHeight || k == len(treeRows)-1 {
					totalScore *= k - i
					break
				}
			}

			if totalScore > tallestTree {
				tallestTree = totalScore
			}
		}
	}

	log.Println(tallestTree)
}
