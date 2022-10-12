//go:build ignore
// +build ignore

package sugar

import (
	"fmt"
	"testing"
)

func TestClamp(t *testing.T) {
	clamp := Clamp(2, 3, 5)
	fmt.Printf("clamp:%v", clamp)
}
