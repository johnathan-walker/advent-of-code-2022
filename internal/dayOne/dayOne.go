package dayone

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/johnathan-walker/advent-of-code-2022/data_structures/heap"
)

type CaloricElf struct {
	Calories int `json:"Calories"`
	ID       int `json:"ID"`
}

func (ce CaloricElf) Value() int {
	return ce.Calories
}

func (ce *CaloricElf) String() string {
	b, err := json.Marshal(ce)
	if err != nil {
		return ""
	}
	return string(b)
}

func RunStarOne(r io.Reader) (int, error) {
	s, h := setupMax(r)
	readIn(s, &h)
	maxElf := h.Pop()
	return maxElf.Calories, nil
}

func RunStarTwo(r io.Reader) ([]int, error) {
	s, h := setupMax(r)
	readIn(s, &h)
	return []int{h.Pop().Calories, h.Pop().Calories, h.Pop().Calories}, nil
}

func setupMax(r io.Reader) (*bufio.Scanner, heap.MaxHeap[CaloricElf]) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	h := make(heap.MaxHeap[CaloricElf], 0)
	h.Init()

	return s, h
}

func readIn(s *bufio.Scanner, h *heap.MaxHeap[CaloricElf]) error {
	i := 1
	currentElf := &CaloricElf{
		Calories: 0,
		ID:       i,
	}

	for s.Scan() {
		token := strings.TrimSpace(s.Text())
		if token == "" {
			h.Push(*currentElf)
			i++
			currentElf = &CaloricElf{
				ID:       i,
				Calories: 0,
			}

			continue
		}

		intToken, err := strconv.Atoi(token)
		if err != nil {
			return fmt.Errorf("could not int parse token %s: %w", token, err)
		}

		currentElf.Calories += intToken
	}

	return nil
}
