//go:build ignore
// +build ignore

package sugar

import (
	"fmt"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	in := make(chan any)
	out := make(chan any, 5)
	rb := NewRingBuffer(in, out)
	go rb.Run()

	for i := 0; i < 10; i++ {
		in <- i
	}

	close(in)

	for res := range out {
		fmt.Println(res)
	}
}
