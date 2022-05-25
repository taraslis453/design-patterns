package strategy

import (
	"fmt"
	"sort"
)

type SortStrategy interface {
	Sort(data []int) []int
}

type DescendingSortStrategy struct{}

func (d DescendingSortStrategy) Sort(data []int) []int {
	// sort in descending order ints
	sort.Ints(data)
	fmt.Println("Sorting in descending order", data)
	return data
}

type AscendingSortStrategy struct{}

func (a AscendingSortStrategy) Sort(data []int) []int {
	// sort in ascending order ints
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	fmt.Println("Sorting in ascending order", data)
	return data
}

// Sort is a context object that holds a strategy object
type Sort struct {
	strategy SortStrategy
}

func (s *Sort) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}
