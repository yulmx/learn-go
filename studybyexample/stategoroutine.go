package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOP struct {
	key int
	resp chan int
}

type writeOP struct {
	key int
	val int
	resp chan bool
}

func main() {

	var ops int64

	reads := make(chan *readOP)
	writes := make(chan *writeOP)


	go func() {
		var state = make(map[int]int)

		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r:=0;r<100;r++	{
		go func() {
			for {
				read := &readOP{key: rand.Intn(5), resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	for w:=0;w<10;w++	{
		go func() {
			for {
				write := &writeOP{key: rand.Intn(5),val: rand.Intn(100), resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops: ", opsFinal)

}