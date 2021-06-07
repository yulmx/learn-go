package main

import (
	"fmt"
	"sort"
)

func main() {

	ss := []string{"c", "a", "b"}
	sort.Strings(ss)
	fmt.Println("Strings:", ss)


	is := []int{7,2,4}
	sort.Ints(is)
	fmt.Println("Ints:", is)

	b := sort.IntsAreSorted(is)
	fmt.Println("Issorted:", b)

}
