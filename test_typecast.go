package main

import (
	"fmt"
	"strconv"
)

type TA int
type TB int
type TC string
type Flyable interface {
	Fly()
}

func (a TA) Fly() {
	fmt.Println("TA.Fly")
}

func main() {
	a := TA(10)
	var b TB = TB(a)
	fmt.Println(b)
	a.Fly()

	var f Flyable = a
	a, ok := f.(TA)
	fmt.Println(a, ok)

	var i Flyable
	ta, ok := i.(TA)
	fmt.Println(ta, ok)

	var s string = strconv.Itoa(1)
	fmt.Println(s)
}
