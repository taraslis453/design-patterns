package strategy

import (
	"fmt"
	"sort"
)

type sortAlgo interface {
	sort(data []int) []int
}

type DescendingSort struct{}

func (d DescendingSort) sort(data []int) []int {
	// sort in descending order ints
	sort.Ints(data)
	fmt.Println("Sorting in descending order", data)
	return data
}

type AscendingSort struct{}

func (a AscendingSort) sort(data []int) []int {
	// sort in ascending order ints
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	fmt.Println("Sorting in ascending order", data)
	return data
}

// Sort is a context object that holds a strategy object
type Sort struct {
	strategy sortAlgo
}

func (s *Sort) setStrategy(strategy sortAlgo) {
	s.strategy = strategy
}
