package main

import "fmt"

func Traverse(m map[int]string) {
	fmt.Println("__________________________")
	for k, v := range m {
		fmt.Printf("[%v]:[%v]\n", k, v)
	}
}

func TestMap() {
	// 创建1
	var m0 = map[string]string{
		"123": "456",
	}
	fmt.Println(m0)

	// 创建2
	var m = make(map[int]string)
	m[1] = "abc"
	m[2] = "123"
	m[3] = "3v4"
	Traverse(m)

	// 添加/修改/查询/删除
	m[4] = "p1v"
	m[3] = "2pv"
	Traverse(m)
	v1 := m[3]
	v2, ok2 := m[3]
	v3, ok3 := m[5]
	fmt.Println(v1)
	fmt.Println(v2, ok2)
	fmt.Println(v3, ok3)
	delete(m, 2)
	Traverse(m)
}

func main() {
	TestMap()
}
