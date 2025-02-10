package main

import (
	"fmt"
	"tidy/interface/dog"
)

type Animal interface {
	Speak() string
}

func speakAnimal(animal Animal) {
	fmt.Println(animal.Speak())
}

func main() {
	cat := &dog.Dog{Name: "Whiskers", Age: 2}
	speakAnimal(cat)
}
