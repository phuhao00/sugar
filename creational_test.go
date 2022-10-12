//go:build ignore
// +build ignore

package sugar

import (
	"fmt"
	"testing"
)

type BuilderImplCase struct {
	M1 any
}

func (b *BuilderImplCase) Build(m any) Builder {
	b.M1 = m
	return b
}

func (b *BuilderImplCase) CompletedCheck() bool {
	if !IsNil1(b.M1) {
		return true
	}
	return false
}

func TestBuilderImplCase(t *testing.T) {
	v := &BuilderImplCase{}
	completedCheck := v.Build("m").CompletedCheck()
	if completedCheck {
		fmt.Println("complete!")
	}
}
