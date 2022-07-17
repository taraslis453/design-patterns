package oop

import (
	"fmt"

	"github.com/taraslis453/design-patterns/oop/inheritance"
	"github.com/taraslis453/design-patterns/oop/polymorphism"
)

func Run() {
	fmt.Println("Running OOP")
	fmt.Println("Polymorphism")
	polymorphism.RunPolymorphism()
	fmt.Println("Inheritance")
	inheritance.RunInheritance()
}
