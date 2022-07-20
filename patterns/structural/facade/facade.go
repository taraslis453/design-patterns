package facade

type ComplexFacade struct {
}

func (c *ComplexFacade) BuildBasement() {
	// Build basement
}

func (c *ComplexFacade) BuildStairs() {
	// Build stairs
}

func (c *ComplexFacade) BuildRoof() {
	// Build roof
}

func (c *ComplexFacade) BuildHouse() {
	c.BuildBasement()
	c.BuildStairs()
	c.BuildRoof()
}

func RunFacade() {
	// Create a complex system to build a house
	complex := &ComplexFacade{}
	complex.BuildHouse()
}
