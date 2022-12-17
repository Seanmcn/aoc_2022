package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"sort"
	"strings"
)

//go:embed input.txt
var file embed.FS

func getItemPriority(item string) int {
	// convert item to rune, as cannot go string -> int
	runes := []rune(item)
	// convert rune back to int, we know it's only one character
	asciiCode := int(runes[0])

	// Map the ascii code to the priority we want
	var priority = 0
	if asciiCode > 90 {
		priority = asciiCode - 96
	} else {
		priority = asciiCode - 65 + 27
	}
	return priority
}

func getDuplicateItem(compartmentA string, compartmentB string) string {
	for _, character := range compartmentA {
		if strings.Contains(compartmentB, string(character)) {
			return string(character)
		}
	}
	return ""
}

func getSharedItem(groupStrings [3]string, allGroupItemCount map[string]int) string {
	for _, s := range groupStrings {
		// Map to ensure we aren't counting the same character twice
		var groupItemCount = map[string]int{}

		for _, c := range s {
			if _, ok := groupItemCount[string(c)]; ok {
				// Already have character, move on
			} else {
				groupItemCount[string(c)] = 1
				allGroupItemCount[string(c)] += 1
			}
		}
	}

	// Make slice of keys to sort
	keys := make([]string, 0, len(allGroupItemCount))
	for k := range allGroupItemCount {
		keys = append(keys, k)
	}

	// Sort by value
	sort.Slice(keys, func(i, j int) bool { return allGroupItemCount[keys[i]] > allGroupItemCount[keys[j]] })

	// Return the highest value
	for _, k := range keys {
		return k
	}

	// Failsafe
	return ""
}

func main() {

	f, err := file.Open("input.txt")
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	sc := bufio.NewScanner(f)
	itemPrioritySum := 0
	groupItemPrioritySum := 0

	c := 1
	var groupStrings [3]string
	var allGroupItemCount = map[string]int{}

	for sc.Scan() {
		line := sc.Text()
		halfLineLen := len(line) / 2

		compartmentA := line[0:halfLineLen]
		compartmentB := line[halfLineLen:]
		duplicateItem := getDuplicateItem(compartmentA, compartmentB)
		itemPriority := getItemPriority(duplicateItem)
		itemPrioritySum += itemPriority

		k := c - 1
		groupStrings[k] = line

		if c%3 == 0 {
			c = 1
			sharedItem := getSharedItem(groupStrings, allGroupItemCount)
			sharedItemPriority := getItemPriority(sharedItem)
			groupItemPrioritySum += sharedItemPriority
			// Reset Map
			allGroupItemCount = map[string]int{}

		} else {
			c++
		}

	}

	fmt.Printf("Part One: : %d\n", itemPrioritySum)

	fmt.Printf("Part Two: : %d \n", groupItemPrioritySum)

}
