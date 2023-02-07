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

var totatThresholdDirSizes int

func part1() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	currentDir := &dir{
		name:           "/",
		subDirs:        make(map[string]*dir),
		parent:         nil,
		totalFileSizes: 0,
	}
	root := currentDir

	index := 0
	for sc.Scan() {
		split := strings.Split(sc.Text(), " ")

		switch split[0] {
		case "$":
			switch split[1] {
			case "cd":
				// Change directory
				//log.Println(currentDir.name, "=>", split[2])

				switch split[2] {
				case "..":
					// Go up one directory
					currentDir = currentDir.parent
				case "/":
					// Go to root
					currentDir = root
				default:
					// Go to directory
					if _, ok := currentDir.subDirs[split[2]]; !ok {
						currentDir.subDirs[split[2]] = &dir{
							name:           split[2],
							subDirs:        make(map[string]*dir),
							parent:         currentDir,
							totalFileSizes: 0,
						}
					}

					currentDir = currentDir.subDirs[split[2]]
				}
			}
			// Command
		case "dir":
			// Directory
			if _, ok := currentDir.subDirs[split[1]]; !ok {
				currentDir.subDirs[split[1]] = &dir{
					name:           split[1],
					subDirs:        make(map[string]*dir),
					parent:         currentDir,
					totalFileSizes: 0,
				}
			}
		default:
			// File listing
			size, err := strconv.Atoi(split[0])
			if err != nil {
				log.Fatal(err)
			}

			currentDir.totalFileSizes += size
		}

		index++
	}

	// TODO: Do we iterate over a directory multiple times?
	// I'm pretty sure that we are not, because the answer was correct
	getTotalBelowThreshold(root, 100000)

	log.Println(totatThresholdDirSizes)
}

var totalDiskSpace int = 70000000
var requiredDiskSpace int = 30000000

func part2() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	currentDir := &dir{
		name:           "/",
		subDirs:        make(map[string]*dir),
		parent:         nil,
		totalFileSizes: 0,
	}
	root := currentDir

	index := 0
	for sc.Scan() {
		split := strings.Split(sc.Text(), " ")

		switch split[0] {
		case "$":
			switch split[1] {
			case "cd":
				// Change directory
				//log.Println(currentDir.name, "=>", split[2])

				switch split[2] {
				case "..":
					// Go up one directory
					currentDir = currentDir.parent
				case "/":
					// Go to root
					currentDir = root
				default:
					// Go to directory
					if _, ok := currentDir.subDirs[split[2]]; !ok {
						currentDir.subDirs[split[2]] = &dir{
							name:           split[2],
							subDirs:        make(map[string]*dir),
							parent:         currentDir,
							totalFileSizes: 0,
						}
					}

					currentDir = currentDir.subDirs[split[2]]
				}
			}
			// Command
		case "dir":
			// Directory
			if _, ok := currentDir.subDirs[split[1]]; !ok {
				currentDir.subDirs[split[1]] = &dir{
					name:           split[1],
					subDirs:        make(map[string]*dir),
					parent:         currentDir,
					totalFileSizes: 0,
				}
			}
		default:
			// File listing
			size, err := strconv.Atoi(split[0])
			if err != nil {
				log.Fatal(err)
			}

			currentDir.totalFileSizes += size
		}

		index++
	}

	// TODO: Do we iterate over a directory multiple times?
	// I'm pretty sure that we are not, because the answer was correct
	totalUsed := getTotalBelowThreshold(root, 100000)

	available := totalDiskSpace - totalUsed
	delta := requiredDiskSpace - available

	// TODO: Could we do this without iterating over the tree again?
	log.Println(getClosestDir(root, delta, root).totalSubDirSizes)
}

type dir struct {
	name             string
	subDirs          map[string]*dir
	parent           *dir
	totalFileSizes   int
	totalSubDirSizes int
}

func getTotalBelowThreshold(dir *dir, threshold int) int {
	total := dir.totalFileSizes

	for _, subDir := range dir.subDirs {
		total += getTotalBelowThreshold(subDir, threshold)
	}

	if total < threshold {
		totatThresholdDirSizes += total
	}

	dir.totalSubDirSizes = total

	return total
}

func getClosestDir(dir *dir, delta int, closestDir *dir) *dir {
	for _, subDir := range dir.subDirs {
		subClosestDir := getClosestDir(subDir, delta, closestDir)
		if subClosestDir.totalSubDirSizes < closestDir.totalSubDirSizes &&
			subClosestDir.totalSubDirSizes >= delta {
			closestDir = subClosestDir
		}

		if subDir.totalSubDirSizes < closestDir.totalSubDirSizes &&
			subDir.totalSubDirSizes >= delta {
			closestDir = subDir
		}
	}

	return closestDir
}
