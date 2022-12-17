package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
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

func main() {

	f, err := file.Open("input.txt")
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	sc := bufio.NewScanner(f)
	itemPrioritySum := 0

	c := 1

	for sc.Scan() {
		line := sc.Text()
		halfLineLen := len(line) / 2

		compartmentA := line[0:halfLineLen]
		compartmentB := line[halfLineLen:]
		duplicateItem := getDuplicateItem(compartmentA, compartmentB)
		itemPriority := getItemPriority(duplicateItem)
		itemPrioritySum += itemPriority

		if c%3 == 0 {
			// new group
		}

		c++

	}

	fmt.Printf("Part One: : %d\n", itemPrioritySum)

	fmt.Printf("Part Two: : N/A")

}
