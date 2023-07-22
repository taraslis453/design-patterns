package abstractfactory

import "fmt"

type iShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shoe struct {
	logo string
	size int
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) getSize() int {
	return s.size
}

type iShort interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type short struct {
	logo string
	size int
}

func (s *short) setLogo(logo string) {
	s.logo = logo
}

func (s *short) setSize(size int) {
	s.size = size
}

func (s *short) getLogo() string {
	return s.logo
}

func (s *short) getSize() int {
	return s.size
}

type brandAShoe struct {
	shoe
}

type brandAShort struct {
	short
}

type brandA struct {
}

func (a *brandA) makeShoe() iShoe {
	return &brandAShoe{
		shoe: shoe{
			logo: "brandA",
			size: 14,
		},
	}
}

func (a *brandA) makeShort() iShort {
	return &brandAShort{
		short: short{
			logo: "brandA",
			size: 14,
		},
	}
}

type brandBShoe struct {
	shoe
}

type brandBShort struct {
	short
}

type brandB struct {
}

func (b *brandB) makeShoe() iShoe {
	return &brandBShoe{
		shoe: shoe{
			logo: "brandB",
			size: 14,
		},
	}
}

func (b *brandB) makeShort() iShort {
	return &brandBShort{
		short: short{
			logo: "brandB",
			size: 14,
		},
	}
}

type iSportsAttireFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsAttireFactory(brand string) (iSportsAttireFactory, error) {
	if brand == "brandA" {
		return &brandA{}, nil
	}
	if brand == "brandB" {
		return &brandB{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}

func RunAbstractFactory() {
	brandAFactory, _ := getSportsAttireFactory("brandA")
	brandBFactory, _ := getSportsAttireFactory("brandB")
	brandBShoe := brandBFactory.makeShoe()
	brandBShort := brandBFactory.makeShort()
	brandAShoe := brandAFactory.makeShoe()
	brandAShort := brandAFactory.makeShort()
	printShoeDetails(brandBShoe)
	printShortDetails(brandBShort)
	printShoeDetails(brandAShoe)
	printShortDetails(brandAShort)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShortDetails(s iShort) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
