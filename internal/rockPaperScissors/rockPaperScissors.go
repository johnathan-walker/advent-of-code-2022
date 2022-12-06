package rockpaperscissors

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func PartOne(r io.Reader) (int, error) {
	return getScore(r, analyzedRound)
}

func PartTwo(r io.Reader) (int, error) {
	return getScore(r, calculateRound)
}

func getScore(r io.Reader, compFunc func([]string) int) (int, error) {
	totalScore := 0
	for s := bufio.NewScanner(r); s.Scan(); {
		moves, err := getMoves(s)
		if err != nil {
			return totalScore, err
		}

		totalScore += compFunc(moves)
	}
	return totalScore, nil
}

func getMoves(s *bufio.Scanner) ([]string, error) {
	line := s.Text()
	moves := strings.Split(line, " ")

	if len(moves) != 2 {
		return nil, errors.New("read a line with more than 2 tokens")
	}

	return moves, nil
}

func analyzedRound(moves []string) int {
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

func calculateRound(moves []string) int {
	score := 0
	return score
}
