package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{x: 1,y: 2}

	fmt.Printf("%%v  %v\n", p)
	fmt.Printf("%%+v %+v\n",p)
	fmt.Printf("%%#v %#v\n", p)
	fmt.Printf("%%T  %T\n", p)
	fmt.Printf("%%t	%t\n", true)
	fmt.Printf("%%d  %d\n", 123)
	fmt.Printf("%%b  %b\n", 123)
	fmt.Printf("%%c  %c\n", 123)
	fmt.Printf("%%x  %x\n", 123)
	fmt.Printf("%%f  %f\n", 12.3)
	fmt.Printf("%%e  %e\n", 1230000.000)
	fmt.Printf("%%s  %s\n", "\"string\"")
	fmt.Printf("%%q  %q\n", "\"string\"")
	fmt.Printf("%%x  %x\n", "\"string\"")
	fmt.Printf("%%p  %p\n", &p)

	fmt.Printf("|%%6d|%%6d|\n")
	fmt.Printf("|%6d|%6d|\n", 123, 23)

	fmt.Printf("|%%6.2f|%%6.2f|\n")
	fmt.Printf("|%6.2f|%6.2f|\n", 1.23, 2131.00)

	fmt.Printf("|%%-6.2f|%%-6.2f|\n")
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.23, 2131.00)

	fmt.Printf("|%%6s|%%6s|\n")
	fmt.Printf("|%6s|%6s|\n", "foo", "a")

	fmt.Printf("|%%-6s|%%-6s|\n")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "a")


	s := fmt.Sprintf("a %s", "string")
	fmt.Println("sprintf:", s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")


}
