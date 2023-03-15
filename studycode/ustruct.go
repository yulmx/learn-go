// @program: learn-go
// @author: yulm
// @created: 2023-03-14 18:37
// @note:
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name  string
	age   int    // 小写私有成员(对外不可见)
	Class string // 首字母大写则该成员为公有成员(对外可见)
	Tag   string "the tag of struct"
}

// 工厂模式自定义构造函数
func NewStudent(name string, age int, class string) *Student {
	return &Student{name: name, age: age, Class: class}
}

// 内存分布
type MemAlloc struct {
	Name string
	Age  int
	w    int64
	h    int8
	s    float64
}

func TestMemoryAlloc() {
	var ma = new(MemAlloc)
	ma.Name = "中国加油"
	fmt.Println("----------------")
	fmt.Printf("%p\n", &ma.Name)
	fmt.Printf("%p\n", &ma.Age)
	fmt.Printf("%p\n", &ma.w)
	fmt.Printf("%p\n", &ma.h)
	fmt.Printf("%p\n", &ma.s)
	typ := reflect.TypeOf(MemAlloc{})
	fmt.Printf("struct MemAlloc has %d bytes\n", typ.Size())
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s offset %v, size=%d, align=%d\n", field.Name, field.Offset, field.Type.Size(), field.Type.Align())
	}
}

func main() {
	var stu1 Student
	var stu2 *Student = new(Student)
	var stu3 *Student = &Student{
		name:  "rose",
		age:   10,
		Class: "class3",
	}

	stu1.name = "china"
	stu1.age = 20
	stu1.Class = "class1"
	fmt.Println(stu1.name)

	stu2.name = "US"
	stu2.age = 24
	fmt.Println(stu2.name, (*stu2).name)

	fmt.Println(stu3.name, (*stu3).name)

	stu4 := NewStudent("May", 34, "class5")
	fmt.Println(stu4.name)

	TestMemoryAlloc()
}
