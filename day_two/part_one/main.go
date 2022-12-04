package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const ScoreResultLoss = 0
const ScoreResultDraw = 3
const ScoreResultWin = 6

func shapeToScore(shape string) int {
	mapping := map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}
	return mapping[shape]
}

func encryptedToValue(value string) string {
	mapping := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}
	return mapping[value]
}

func gameResult(p1Action string, p2Action string) string {
	winConditions := map[string]string{
		"rock":     "scissors",
		"scissors": "paper",
		"paper":    "rock",
	}
	playerOneWin := winConditions[p1Action]
	if p1Action == p2Action {
		return "draw"
	} else if p2Action == playerOneWin {
		return "win"
	} else {
		return "loss"
	}
}

func main() {
	f, err := os.OpenFile("example.txt", os.O_RDONLY, os.ModePerm)
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

	sc := bufio.NewScanner(f)
	score := 0
	for sc.Scan() {
		line := sc.Text()
		moves := strings.Fields(line)
		result := gameResult(encryptedToValue(moves[1]), encryptedToValue(moves[0]))

		switch result {
		case "win":
			score += shapeToScore(encryptedToValue(moves[1])) + ScoreResultWin
		case "draw":
			score += shapeToScore(encryptedToValue(moves[1])) + ScoreResultDraw
		case "loss":
			score += shapeToScore(encryptedToValue(moves[1])) + ScoreResultLoss
		}
	}

	fmt.Printf("Score is: %d\n", score)

}
