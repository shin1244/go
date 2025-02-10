package main

import (
	"fmt"
	"tidy/interface/duck"
)

type Animal interface {
	Speak() string
}

func speakAnimal(animal Animal) {
	fmt.Println(animal.Speak())
}

func main() {
	cat := &duck.Duck{Name: "Whiskers", Age: 2}
	speakAnimal(cat)
}
