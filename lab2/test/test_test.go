package test

import (
	"testing"
)

func a() func() int {
	data := 1
	return func() int {
		data++
		return data
	}
}
func TestMainTest(t *testing.T) {
	saved := a()
	for i := 0; i < 500; i++ {
		t.Log(saved())
	}
}
