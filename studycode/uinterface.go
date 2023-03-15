package main

import "fmt"

// 定义接口
type Skills interface {
    Getname() string
    Running()
}

// 接口嵌套
type Skills1 interface {
    Sleeping()
    Skills
}

// 接口组合
type Playable interface {
    Playing()
}

type AllSkills interface {
    Skills
    Playable
}

type Student struct {
    Name string
    Age int
}

// 实现接口
func (s Student) Getname() string {
    fmt.Println(s.Name)
    return s.Name
}

func (s Student) Running() {
    fmt.Printf("%s running\n", s.Name)
}

func (s Student) Sleeping() {
    fmt.Printf("%s sleeping\n", s.Name)
}

func (s Student) Playing() {
    fmt.Printf("%s playing\n", s.Name)
}

// 使用接口：思考这么写的意义在哪里？？？
func Test1() {
    var skill Skills
    var stu Student
    stu.Age = 20
    stu.Name = "Bob"
    skill = stu
    skill.Running()     // 调用接口
}

func Test2() {
    var skill Skills1
    var stu Student
    stu.Age = 20
    stu.Name = "Bob"
    skill = stu
    skill.Running()     // 调用接口
    skill.Sleeping()
}

func Test3() {
    var skill AllSkills
    var stu Student
    stu.Age = 20
    stu.Name = "Bob"
    skill = stu
    skill.Running()     // 调用接口
    skill.Playing()
}

// 类型转换
func Test4() {
    var i int
    var x interface{}

    x = i   // 为什么能赋值到空接口, 每种类型都已经隐藏实现了空接口
    y, ok := x.(int)  ////将interface 转为int,ok可省略 但是省略以后转换失败会报错，true转换成功，false转换失败, 并采用默认值
    z, ok1 := x.(string)
    fmt.Println(y, ok)
    fmt.Println(z, ok1)
}

// 类型转换
func TestType(items ...interface{}) {
    for k, v := range items {
        switch v.(type) {
            case string:
            fmt.Printf("type is string, %d: %v\n", k, v)
            case bool:
            fmt.Printf("type is bool, %d: %v\n", k, v)
            case int:
            fmt.Printf("type is int, %d: %v\n", k, v)
            case float32, float64:
            fmt.Printf("type is float, %d: %v\n", k, v)
            case Student:
            fmt.Printf("type is struct Student, %d: %v\n", k, v)
            case *Student:
            fmt.Printf("type is struct pointer Student, %d: %p\n", k, v)
        }
    }
}

func Test5() {
    var stu Student
    TestType("King", 35, 3.33, stu, true)
}

// 接口变量
func Test6() {
    var a interface{}
    var s = Student{"Joe", 25}
    a = &s          // 指向指针

    var ps = a.(*Student)   // 转换成指针类型
    s.Name = "AKI"
    s.Age = 21
    fmt.Println("s: ", s)
    fmt.Println("ps: ", ps)
    fmt.Printf("ps: %p, s: %p\n", ps, &s)   // 输出结构 ps = &s
}

func main() {
    // Test1()
    // Test2()
    // Test3()
    Test4()
    // Test5()
    // Test6()
}