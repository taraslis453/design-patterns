package polymorphism

import "fmt"

type income interface {
	calculate() int
	source() string
}

type company struct {
	name        string
	employees   []employee
	totalIncome int
}

type employee struct {
	firstName   string
	lastName    string
	hoursWorked float64
}

func (c *company) calculate() int {
	for _, employee := range c.employees {
		c.totalIncome += int(employee.hoursWorked * 150)
	}
	return c.totalIncome
}

func (c *company) source() string {
	return c.name
}

type client struct {
	firstName string
	lastName  string
	totalPaid int
}

func (c *client) calculate() int {
	return c.totalPaid
}

func (c *client) source() string {
	return c.firstName + " " + c.lastName
}

func getIncomeInvoices(invoices []income) {
	for _, invoice := range invoices {
		fmt.Printf("%s earned %d\n", invoice.source(), invoice.calculate())
	}
}

func RunPolymorphism() {
	// Create two companies, and two clients.
	company1 := company{
		name: "Company One",
		employees: []employee{
			{
				firstName:   "Joe",
				lastName:    "Bloggs",
				hoursWorked: 10.5,
			},
			{
				firstName:   "Jane",
				lastName:    "Smith",
				hoursWorked: 12.5,
			},
		},
	}
	company2 := company{
		name: "Company Two",
		employees: []employee{
			{
				firstName:   "John",
				lastName:    "Doe",
				hoursWorked: 5.5,
			},
			{
				firstName:   "Jane",
				lastName:    "Doe",
				hoursWorked: 11.5,
			},
		},
	}
	client1 := client{
		firstName: "John",
		lastName:  "Doe",
		totalPaid: 100,
	}
	client2 := client{
		firstName: "Jane",
		lastName:  "Doe",
		totalPaid: 200,
	}

	// Add the companies and clients to a slice.
	items := []income{&company1, &company2, &client1, &client2}

	getIncomeInvoices(items)
}
