package main

import (
	"fmt"
	"os"
)

type opfn[T ~int | ~float64] func(i ...T) T
type writer func(string)

func writeHello(f writer) {
	f("Hello, World!")
}

func sum[T ~int | ~float64](i ...T) T {
	var sum T
	for _, v := range i {
		sum += v
	}
	return sum
}

func mul[T ~int | ~float64](i ...T) T {
	var mul T
	mul = 1
	for _, v := range i {
		mul *= v
	}
	return mul
}

func operator[T ~int | ~float64](op string) opfn[T] {
	if op == "+" {
		return sum[T]
	} else if op == "*" {
		return mul[T]
	}
	return nil
}

func create() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer fmt.Println("파일 닫기")

	fmt.Fprintln(f, "Hello, World!")
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(100, 200))
	fmt.Println(sum(0.3, 0.7))
	create()
	op := operator[int]("+")
	fmt.Println(op(1, 2, 3, 4, 5))
	op = operator[int]("*")
	fmt.Println(op(1, 2))
	writeHello(func(s string) {
		fmt.Println(s)
	})
}
