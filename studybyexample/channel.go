package main

import (
	"fmt"
	"time"
)

func main()	{

	message := make(chan string)

	go func() {message <- "ping"}()

	msg := <-message
	fmt.Println(msg)


	// 通道缓冲
	message2 := make(chan string, 2)
	message2 <- "buffered"
	message2 <- "channel"

	fmt.Println(<-message2)
	fmt.Println(<-message2)

	//message <- "test"			// fatal error: all goroutines are asleep - deadlock!
	//fmt.Println(<-message)


	// 通道同步
	done := make(chan bool, 1)
	go worker(done)

	//<-done	// 未进行worker就结束了


	// 通道方向
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "pass message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func worker(done chan bool)  {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true

}

// 通道方向
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string)  {
	msg := <-pings
	pongs <- msg
}
