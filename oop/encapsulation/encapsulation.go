// Encapsulation is a package that encapsulates the data and methods of a person.
package encapsulation

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p *Person) SetFirstName(firstName string) {
	p.printInfo()
	p.firstName = firstName
}

func (p *Person) SetLastName(lastName string) {
	p.printInfo()
	p.lastName = lastName
}

func (p *Person) SetAge(age int) {
	p.printInfo()
	p.age = age
}

func (p *Person) printInfo() {
	fmt.Printf("%s %s is %d years old.\n", p.firstName, p.lastName, p.age)
}
