package main

import "fmt"

func test(from string) {
	for i:=0; i<10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	test("direct")

	go test("goroutines")

	go func(msg string) {
		fmt.Println(msg)
	}("goingabcdef")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}
