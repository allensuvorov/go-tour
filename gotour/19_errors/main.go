/*
type error interface {
	Error() string
}
*/

/*
i, err := strconv.Atoi("42")
if err != nil {
	fmt.Printf("couldn't convert number: %v\n, err)
	return
}
fmt.Println("Converted integer:", i)
*/

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it did't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
