package main

import "fmt"

type name struct {
	first string
	last  string
}

type user struct {
	n   *name
	age int
}

func main() {
	u := user{age: 18}
	fmt.Println("first", u.n.first)
	fmt.Println("u.n", u.n == nil)
	u.n.first = "zzp"
	fmt.Println("first", u.n.first)
}
