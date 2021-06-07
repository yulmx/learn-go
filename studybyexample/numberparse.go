package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	f1, f1e := strconv.ParseFloat("1.234bbad12", 64)
	fmt.Println(f1, f1e)

	i, _ := strconv.ParseInt("1234", 0, 64)
	fmt.Println(i)
	i1, i1e := strconv.ParseInt("12dfs34", 0, 64)
	fmt.Println(i1, i1e)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseInt("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("12532")
	fmt.Println(k)
	_, e := strconv.Atoi("wait")
	fmt.Println(e)

}