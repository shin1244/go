package cat

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func (c *Cat) Speak() string {
	return fmt.Sprintf("Meow! My name is %s and I am %d years old", c.Name, c.Age)
}
