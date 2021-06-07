// 命令行标志


package main

import (
	"flag"
	"fmt"
)

func main() {

	word := flag.String("word", "foo", "a string")
	number := flag.Int("number", 42, "an int")
	fork := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word\t", *word)
	fmt.Println("number\t", *number)
	fmt.Println("fork\t", *fork)
	fmt.Println("svar\t", svar)
	fmt.Println("tail\t", flag.Args())

}