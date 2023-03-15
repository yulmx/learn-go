// @program: learn-go
// @author: yulm
// @created: 2023-03-15 09:40
// @note: Goroutine
package main

import (
	"fmt"
	"runtime"
	"time"
)

func hello() {
	fmt.Println("goroutine running")
}

func testChildGoroutine() {
	fmt.Println("main goroutine")
	go func() {
		fmt.Println("child goroutine")
		go func() {
			fmt.Println("grant child goroutine")
			go func() {
				// 添加异常处理-----------------------
				defer func() {
					fmt.Println("run defer------")
					if err := recover(); err != nil {
						fmt.Println("捕获error: ", err)
					}
				}()
				// -----------------------添加异常处理
				fmt.Println("grant grant child goroutine")
				var ptr *int
				*ptr = 0x12345 // 故意制作崩溃
			}()
		}()
	}()
	time.Sleep(time.Second)
	fmt.Println("main goroutine will quit")
}

// 同时开启百万协程
func testMillionRoutines() {
	const N = 1000000
	i := 1
	for {
		go func() {
			for {
				time.Sleep(time.Second)
			}
		}()

		if i%10000 == 0 {
			fmt.Printf("%d goroutine started!\n", i)
		}

		i++
		if i == N {
			break
		}
	}
}

// 死循环：其他协程继续执行
func testEndlessLoop() {
	n := 3
	for i := 0; i < n; i++ {
		go func() {
			fmt.Println("dead loop goroutine start")
			for {
			} // 死循环
		}()
	}
	for {
		time.Sleep(time.Second)
		fmt.Println("other goroutine running")
	}
}

// 线程数获取与设置
func testSetThreads(n int) {
	// 读取默认线程数
	fmt.Println(runtime.GOMAXPROCS(0))

	// 设置线程数
	runtime.GOMAXPROCS(n)
	fmt.Println(runtime.GOMAXPROCS(0))
}

func main() {
	//go hello()
	//
	//time.Sleep(3 * time.Second)
	//fmt.Println("main running")

	//testChildGoroutine()
	//testExcept()

	//fmt.Println("run in main")
	//testMillionRoutines()
	//time.Sleep(time.Second * 10)

	//testEndlessLoop()

	testSetThreads(3)
}
