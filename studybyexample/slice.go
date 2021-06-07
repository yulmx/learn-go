package main

import "fmt"
// slice 即 list
func main () {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "123"
	s[2] = "AB"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[1])

	fmt.Println("len(slice): ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl3: ", l)

	t := []string{"3123", "afa", "ABC"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
	fmt.Println("emp: ", twoD)
	for i:=0; i < 3; i++ {
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j:=0; j<innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
