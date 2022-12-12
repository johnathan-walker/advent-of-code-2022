package rucksack

import (
	"bufio"
	"fmt"
	"io"
)

func PartOne(f io.Reader) (int, error) {
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	input := []string{}
	for s.Scan() {
		input = append(input, s.Text())
	}

	return scoreSaks(input)
}

func scoreSaks(sacks []string) (int, error) {
	if len(sacks) == 0 {
		return 0, nil
	}

	total := 0
	for _, sack := range sacks {
		seen := map[rune]int{}
		mid := len(sack) / 2
		compOne := sack[:mid]
		compTwo := sack[mid:]

		fmt.Printf("\nOriginal:\t%s\n", sack)
		fmt.Printf("Compartment 1:\t%s\n", compOne)
		fmt.Printf("Compartment 2:\t%s\n", compTwo)

		for i := range compOne {
			seen[rune(compOne[i])]++
		}

		shouldBreak := false
		for i := range compTwo {
			if _, present := seen[rune(compTwo[i])]; present {
				ascii := int(compTwo[i])
				switch {
				case ascii >= 65 && ascii <= 90:
					score := ascii - 65 + 27
					fmt.Printf("%s scored %d\n", string(compTwo[i]), score)
					total += score
					shouldBreak = true
				case ascii >= 97 && ascii <= 122:
					score := ascii - 96
					fmt.Printf("%s scored %d\n", string(compTwo[i]), score)
					total += score
					shouldBreak = true
				}
			}

			if shouldBreak {
				shouldBreak = false
				break
			}
		}
	}
	return total, nil

}
