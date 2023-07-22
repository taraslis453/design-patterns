package factorymethod

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	name string
}

func NewDog(name string) Dog {
	return Dog{
		name: name,
	}
}

func (d Dog) Speak() string {
	return "woof"
}

type Cat struct {
	name string
}

func NewCat(name string) Cat {
	return Cat{
		name: name,
	}
}

func (c Cat) Speak() string {
	return "Meow"
}

// factory method
func SpeakerFactory(t, name string) Speaker {
	switch {
	case t == "dog":
		return NewDog(name)
	case t == "cat":
		return NewCat(name)
	}
	return nil
}

func FactoryMethod() {
	d := SpeakerFactory("dog", "marshal")
	fmt.Println(d.Speak())
	c := SpeakerFactory("cat", "pi")
	fmt.Println(c.Speak())
}
