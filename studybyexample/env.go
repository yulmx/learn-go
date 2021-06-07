// 环境变量

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println(os.Environ())
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair)
		fmt.Println(pair[0])
	}

}