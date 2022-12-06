package dayone

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func Run(r io.ReadCloser) (int, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	h := make(heap.MaxHeap[CaloricElf], 0)
	h.Init()
	i := 1
	currentElf := &CaloricElf{
		Calories: 0,
		ID:       i,
	}
	for s.Scan() {
		token := strings.TrimSpace(s.Text())
		if token == "" {
			log.Printf("Pushing %s\n", currentElf)
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
			return 0, fmt.Errorf("could not int parse token %s: %w", token, err)
		}

		currentElf.Calories += intToken
	}
	maxElf := h.Pop()
	return maxElf.Calories, nil
}
