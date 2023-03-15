// @program: learn-go
// @author: yulm
// @created: 2023-03-14 22:33
// @note: study method
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) TestMethod() string {
	fmt.Println(p.Name, p.Age)
	return p.Name
}

func (p *Person) TestMethod2() string {
	fmt.Println(p.Name, p.Age)
	return p.Name
}

func main() {
	var p = Person{Name: "Apple", Age: 22}
	var p1 = &Person{Name: "Oracle", Age: 34}
	fmt.Println(p.TestMethod())
	fmt.Println(p1.TestMethod2())
}
