package duck

import "fmt"

type Duck struct {
	Name string
	Age  int
}

func (d *Duck) Speak() string {
	return fmt.Sprintf("Quack! My name is %s and I am %d years old", d.Name, d.Age)
}
