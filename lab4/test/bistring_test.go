package test

import (
	"testing"
	"yanandco/lab4/bitstream"
)

func TestAddBits(t *testing.T) {
	bs := bitstream.NewBitStream()
	bs.WriteBits(0b100111, 6)
	bs.WriteBits(0b1111, 4)

	first := bs.ReadBits(1)
	t.Log(first)
	if first != 0b1 {
		t.Fail()
	}
	second := bs.ReadBits(9)
	t.Log(second)
	if second != 0b111100111 {
		t.Fail()
	}
}
