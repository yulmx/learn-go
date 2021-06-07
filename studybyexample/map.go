package main

import (
	"fmt"
)

// 即dict、hash

func main() {

	m := make(map[string]int)
	fmt.Println("emp: ", m)

	m["k1"] = 12
	m["k2"] = 434
	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("del: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

}
