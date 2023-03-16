// @program: learn-go
// @author: yulm
// @created: 2023-03-16 14:27
// @note: channel
package main

import "fmt"

func testChannel() {
	var bufferedChannel = make(chan int, 1024)
	//var unbufferedChannel = make(chan int)

	for i := 0; i < cap(bufferedChannel); i++ {
		bufferedChannel <- i // 写通道
	}

	for len(bufferedChannel) > 0 {
		var value = <-bufferedChannel // 读通道
		fmt.Println(value)
	}
}

func testChannelBlock() {

}

func testChannelClose() {
	var ch = make(chan int, 4)
	ch <- 1
	ch <- 2
	close(ch)

	value := <-ch
	fmt.Println("1---------", value)
	value = <-ch
	fmt.Println("2---------", value)
	value = <-ch
	fmt.Println("3---------", value)
	//ch <- 3

	var ch1 = make(chan string, 4)
	ch1 <- "abc"
	close(ch1)
	value1 := <-ch1
	fmt.Println("4---------", value1)
}

func main() {
	//testChannel()
	testChannelClose()
}
