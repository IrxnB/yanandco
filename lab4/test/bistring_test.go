package test

import (
	"testing"
	"yanandco/lab4/bitstream"
)

func TestAddBits(t *testing.T) {
	bs := bitstream.NewBitStream()
	bs.WriteBits(0b100111, 6)
	first := bs.ReadBits(2)
	t.Log(first)
	if first != 0b11 {
		t.Fail()
	}
	second := bs.ReadBits(4)
	t.Log(second)
	if second != 0b1001 {
		t.Fail()
	}
}
