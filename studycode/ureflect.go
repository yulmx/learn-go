package main

import (
    "fmt"
    "reflect"
)

// 反射，TypeOf/ValueOf
func Test1() {
    fmt.Println("____________________")
    var s int = 24
    fmt.Println(reflect.TypeOf(s))
    fmt.Println(reflect.ValueOf(s))
}

func Test2() {
    fmt.Println("____________________")
    var s int = 24
    fmt.Println(reflect.TypeOf(s))
}

// Value方法
func Test3() {
    fmt.Println("____________________")
    str := "Golang"
    val := reflect.ValueOf(str).Kind()
    fmt.Println(val)
}

// 修改值
func Test4() {
    fmt.Println("____________________")
    var s int = 55

    // var v = reflect.ValueOf(s)        // 报错 反射值变量
    var v = reflect.ValueOf(&s)         // ok  反射指针变量
    
    v.Elem().SetInt(10)     // 对指针指向的元素进行修改
    fmt.Println(s)
}

func main() {
    Test1()
    Test2()
    Test3()
    Test4()
}
