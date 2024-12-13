package test

import (
	"testing"
	"yanandco/lab4/bitstring"
)

func TestAddBits(t *testing.T) {
	bs := bitstring.NewBitStream()
	bs.WriteBits(0b100111, 6)
	first := bs.ReadBits(2)
	t.Log(first)
	if first != 0b10 {
		t.Fail()
	}
	second := bs.ReadBits(4)
	t.Log(second)
	if second != 0b0111 {
		t.Fail()
	}
}
