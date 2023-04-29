// @program: learn-go
// @author: yulm
// @created: 2023-03-17 19:13
// @note: context
package main

import "time"

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Error() error
	Value(key interface{}) interface{}
}
