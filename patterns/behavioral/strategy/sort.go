package strategy

import (
	"fmt"
	"sort"
)

type SortStrategy interface {
	sort(data []int) []int
}

// Sort is a context object that holds a strategy object
type Sort struct {
	strategy SortStrategy
}

func (s *Sort) setStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

type DescendingSortStrategy struct{}

func (d DescendingSortStrategy) sort(data []int) []int {
	// sort in descending order ints
	sort.Ints(data)
	fmt.Println("Sorting in descending order", data)
	return data
}

type AscendingSortStrategy struct{}

func (a AscendingSortStrategy) sort(data []int) []int {
	// sort in ascending order ints
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	fmt.Println("Sorting in ascending order", data)
	return data
}
