package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.OpenFile("data.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	var currentElf = 1
	elfMap := make(map[int]int)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if len(line) > 0 {
			// Convert line to int
			i, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			// If the elf already exists in the map, update the calorie value for it
			if val, ok := elfMap[currentElf]; ok {
				elfMap[currentElf] = val + i
			} else {
				// Elf doesn't exist in the map, add it
				elfMap[currentElf] = i
			}
		} else {
			// Blank line indicates we are looking at a new elf, increase currentElf value
			currentElf++
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}

	// Create a struct, so we can create a slice to sort
	type ElfCalorie struct {
		Elf      int
		Calories int
	}

	// Add each element to slice
	var sortedElvesByCalories []ElfCalorie
	for k, v := range elfMap {
		sortedElvesByCalories = append(sortedElvesByCalories, ElfCalorie{k, v})
	}

	// Sort the slice by calories
	sort.Slice(sortedElvesByCalories, func(i, j int) bool {
		return sortedElvesByCalories[i].Calories > sortedElvesByCalories[j].Calories
	})

	var topThree = 0
	for c, ec := range sortedElvesByCalories {
		topThree += ec.Calories
		if c == 0 {
			fmt.Printf("Part One Answer: %d\n", ec.Calories)
		}
		if c == 2 {
			fmt.Printf("Part Two Answer: %d\n", topThree)
			break
		}
	}

}
