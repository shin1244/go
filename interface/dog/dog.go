package dog

import "fmt"

type Dog struct {
	Name string
	Age  int
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("Woof! My name is %s and I am %d years old", d.Name, d.Age)
}
