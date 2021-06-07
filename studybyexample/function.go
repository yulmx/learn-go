package main

import "fmt"

// 函数定义
func plus(a int, b int) int {
	return a + b
}

// 返回多值
func vals() (int, int) {
	return 3, 7
}

// 变参函数
func sum(nums ...int){
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 闭包	隐藏变量i
func intSeq() func() int{
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func main() {
	res := plus(1,2)
	fmt.Println("1+2=", res)

	a, b := vals()
	fmt.Println(a,b)

	_, c := vals()
	fmt.Println(c)

	sum(1,2)
	sum(1,2,3)

	nums := []int{1,2,3,4}
	sum(nums...)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newNextInt := intSeq()
	fmt.Println(newNextInt())


	fmt.Println(fact(7))
}