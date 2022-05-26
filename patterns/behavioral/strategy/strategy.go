package strategy

// Strategy is a behavioral design pattern that defines a family of algorithms
// that can be used to solve a particular problem.
func Strategy() {
	var context Sort
	nums := []int{6, 1, 2, 4, 3, 5, 7, 8, 9, 10}
	// easily apply different sorting algorithms based on the input on the fly
	if len(nums) > 4 {
		context.setStrategy(AscendingSort{})
	} else {
		context.setStrategy(DescendingSort{})
	}

	context.strategy.sort(nums)
}
