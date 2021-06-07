package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	now := time.Now()	// 2021-06-07 10:42:23.472363 +0800 CST m=+0.000095403
	p(now)

	then := time.Date(2021, 6, 7, 20, 34, 58, 52345234, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.YearDay())

	p(then.Weekday())
	p(then.ISOWeek())


	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

}