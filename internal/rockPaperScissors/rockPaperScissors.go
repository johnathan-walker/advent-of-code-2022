package rockpaperscissors

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func PartOne(r io.Reader) (int, error) {
	return getScore(r, calcPartOne)
}

func PartTwo(r io.Reader) (int, error) {
	return getScore(r, calcPartTwo)
}

func getScore(r io.Reader, calc func([]string) int) (int, error) {
	totalScore := 0
	for s := bufio.NewScanner(r); s.Scan(); {
		moves, err := getMoves(s)
		if err != nil {
			return totalScore, err
		}

		totalScore += calc(moves)
	}
	return totalScore, nil
}

func getMoves(s *bufio.Scanner) ([]string, error) {
	moves := strings.Split(strings.ToUpper(s.Text()), " ")
	if len(moves) != 2 {
		return nil, errors.New("read a line with more than 2 tokens")
	}
	return moves, nil
}

func calcPartOne(moves []string) int {
	score := 0
	switch moves[1] {
	case "X": // Rock
		score += 1
		switch moves[0] {
		case "A": // Rock
			score += 3
		case "B": // Paper
			score += 0
		case "C": // Scissors
			score += 6
		}
	case "Y": // Paper
		score += 2
		switch moves[0] {
		case "A": // Rock
			score += 6
		case "B": // Paper
			score += 3
		case "C": // Scissors
			score += 0
		}
	case "Z": // Scissors
		score += 3
		switch moves[0] {
		case "A": // Rock
			score += 0
		case "B": // Paper
			score += 6
		case "C": // Scissors
			score += 3
		}
	}
	return score
}

func calcPartTwo(moves []string) int {
	score := 0
	switch moves[1] {
	case "X": // Lose
		score += 0
		switch moves[0] {
		case "A": // Opponent Picks Rock
			score += 3 // Pick Scissors
		case "B": // Opponent Picks Paper
			score += 1 // Pick Rock
		case "C": // Opponent Picks Scissors
			score += 2 // Pick Paper
		}
	case "Y": // Draw
		score += 3
		switch moves[0] {
		case "A": // Opponent Picks Rock
			score += 1 // Pick Rock
		case "B": // Opponent Picks Paper
			score += 2 // Pick Paper
		case "C": // Opponent Picks Scissors
			score += 3 // Pick Scissors
		}
	case "Z": // Win
		score += 6
		switch moves[0] {
		case "A": // Opponent Picks Rock
			score += 2 // Pick Paper
		case "B": // Opponent Picks Paper
			score += 3 // Pick Scissors
		case "C": // Opponent Picks Scissors
			score += 1 // Pick Rock
		}
	}
	return score
}
