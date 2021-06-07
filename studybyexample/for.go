package main

import "fmt"

func main() {
	i := 1

	// 类似C的while
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 类似C中的for
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 无限循环
	for {
		fmt.Println("loop")
		break
	}

}
