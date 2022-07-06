package bridge

import "fmt"

// Abstraction
type computer interface {
	print(string)
	setPrinter(printer printer)
}

var _ computer = &macComputer{}

// Refined Abstraction
type macComputer struct {
	printer printer
}

func (c *macComputer) setPrinter(printer printer) {
	c.printer = printer
}

func (c *macComputer) print(text string) {
	fmt.Println("do some stuff before printing")
	fmt.Println("printing")
	c.printer.print(text)
}

// Implementation
type printer interface {
	print(string)
}

// Concrete Implementation
type samsungPrinter struct {
}

func (p *samsungPrinter) print(text string) {
	fmt.Println("printing with samsung printer: " + text)
}

// Bridge pattern is used to decouple an abstraction from its implementation so that the two can vary independently.
func RunBridge() {
	mac := macComputer{}
	mac.setPrinter(&samsungPrinter{})
	mac.print("hello")
}
