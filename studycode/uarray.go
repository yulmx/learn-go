// @program: learn-go
// @author: yulm
// @created: 2023-03-18 14:59
// @note: 数组array
package main

import "fmt"

func modifyArray(arr [5]int) {
	arr[0] = 100
	return
}

func modifyArrayPtr(arr *[5]int)  {
	(*arr)[0] = 100
	return
}


// 数组是值类型，修改副本不会改变本身
func testArrayModify()  {
	var arr [5]int

	modifyArray(arr)
	fmt.Println("-----传数组-----")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	modifyArrayPtr(&arr)
	fmt.Println("-----传数组指针-----")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	testArrayModify()
}