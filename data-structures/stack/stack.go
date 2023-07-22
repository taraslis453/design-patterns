package stack

import (
	"fmt"
	"log"
)

type stack struct {
	nodes []node
	top   *node
}

type node struct {
	value interface{}
	next  *node
}

func (s *stack) Push(value interface{}) {
	newNode := node{
		value: value,
	}

	if s.top == nil {
		s.top = &newNode
	}

	if len(s.nodes) != 0 {
		s.nodes[len(s.nodes)-1].next = &newNode
	}
	s.nodes = append(s.nodes, newNode)
}

func (s *stack) Pop() (*node, error) {
	size := len(s.nodes)
	if size == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	s.nodes = s.nodes[:size-1]
	return s.top, nil
}

func (s *stack) Peek() (*node, error) {
	if s.nodes == nil {
		fmt.Println("stack is empty")
	}
	return s.top, nil
}

func RunStack() {
	s := stack{}
	s.Push(1)
	s.Push(2)

	fmt.Println(s.nodes[0])

	topNode, err := s.Pop()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(topNode)

	topNode, err = s.Peek()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(topNode)
}
