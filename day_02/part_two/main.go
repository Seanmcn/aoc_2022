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

func resultWeWant(value string) string {
	mapping := map[string]string{
		"X": "lose",
		"Y": "draw",
		"Z": "win",
	}
	return mapping[value]
}

func moveToPlay(resultWeWant string, playersMove string) string {
	winConditions := map[string]string{
		"scissors": "rock",
		"paper":    "scissors",
		"rock":     "paper",
	}
	lossConditions := map[string]string{
		"rock":     "scissors",
		"scissors": "paper",
		"paper":    "rock",
	}
	if resultWeWant == "draw" {
		return playersMove
	} else if resultWeWant == "win" {
		return winConditions[playersMove]
	} else {
		return lossConditions[playersMove]
	}
}

func encryptedToValue(value string) string {
	mapping := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
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
	f, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
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
		theirMove := encryptedToValue(moves[0])
		ourMove := moveToPlay(resultWeWant(moves[1]), theirMove)
		result := gameResult(ourMove, theirMove)

		switch result {
		case "win":
			score += shapeToScore(ourMove) + ScoreResultWin
		case "draw":
			score += shapeToScore(ourMove) + ScoreResultDraw
		case "loss":
			score += shapeToScore(ourMove) + ScoreResultLoss
		}
	}

	fmt.Printf("Final Score is: %d\n", score)

}
