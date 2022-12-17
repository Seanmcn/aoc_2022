package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

//go:embed input.txt
//go:embed example.txt
var file embed.FS

type sectionAssignment struct {
	start int
	end   int
}

func assigmentOverlapsCompletely(first sectionAssignment, second sectionAssignment) bool {
	return (first.start >= second.start && first.end <= second.end) || (second.start >= first.start && second.end <= first.end)
}

func assignmentOverlapsPartially(first sectionAssignment, second sectionAssignment) bool {
	return (first.start <= second.end && first.end >= second.start) || (second.start <= first.end && second.end >= first.start)
}

func main() {
	f, err := file.Open("input.txt")
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}

	sc := bufio.NewScanner(f)

	var re = regexp.MustCompile(`(?m)(\d.*)-(\d.*),(\d.*)-(\d.*)`)

	assignmentsOverlappedCompletely := 0
	assigmentOverlappedPartially := 0
	for sc.Scan() {
		line := sc.Text()
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			elfAStart, _ := strconv.Atoi(match[1])
			elfAEnd, _ := strconv.Atoi(match[2])
			elfBStart, _ := strconv.Atoi(match[3])
			elfBEnd, _ := strconv.Atoi(match[4])

			firstAssignment := sectionAssignment{start: elfAStart, end: elfAEnd}
			secondAssignment := sectionAssignment{start: elfBStart, end: elfBEnd}

			overlapsCompletely := assigmentOverlapsCompletely(firstAssignment, secondAssignment)
			if overlapsCompletely {
				assignmentsOverlappedCompletely++
			}

			overlapsPartially := assignmentOverlapsPartially(firstAssignment, secondAssignment)

			if overlapsPartially {
				assigmentOverlappedPartially++
			}

		}
	}

	fmt.Printf("Answer Part One: %v\n", assignmentsOverlappedCompletely)
	fmt.Printf("Answer Part Two: %v\n", assigmentOverlappedPartially)
}
