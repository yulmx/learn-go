package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 20})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 39})

	s := person{name: "Sean", age: 25}
	fmt.Println(s.name, s.age)

	sp := &s
	fmt.Println(sp.name, sp.age)

	sp.age = 51
	fmt.Println(sp.name, sp.age)
}
