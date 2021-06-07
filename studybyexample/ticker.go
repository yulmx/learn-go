package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())

	ticker := time.NewTicker(time.Second * 3)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 10)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}
