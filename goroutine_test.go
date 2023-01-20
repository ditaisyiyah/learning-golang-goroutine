package goroutine

import (
	"fmt"
	"testing"
	"time"
)

/*
A goroutine is a lightweight thread managed by the Go runtime.

go keyword to make a func to be processed asynchronously
but, only non-void func.
*/

func RunHelloWorld() {
	fmt.Println("Haloo")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // run asynchronously by goroutine
	fmt.Println("Ups")

	time.Sleep(5 * time.Second)
}
