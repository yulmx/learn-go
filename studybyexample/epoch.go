package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	secs := now.Unix()
	nancs := now.UnixNano()
	fmt.Println(now)
	fmt.Println(secs)
	fmt.Println(nancs)

	millis := nancs / 1000000
	fmt.Println(millis)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nancs))

}