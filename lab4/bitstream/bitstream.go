package bitstream

import (
	"yanandco/lab1/crypto"
)

// implementation of a FIFO stream
type BitStream struct {
	str       []bool
	bitLength int
}

func NewBitStream() *BitStream {
	return &BitStream{str: make([]bool, 0), bitLength: 0}
}

func (bs *BitStream) WriteBits(bitsToAdd byte, numBits int) {
	for i := numBits - 1; i >= 0; i-- {
		bs.str = append(bs.str, (bitsToAdd>>i)&1 == 1)
		bs.bitLength++
	}
}

func (bs *BitStream) WriteTelegraphChar(t crypto.TelegraphChar) {
	bs.WriteBits(t.Char, 5)
}

func (bs *BitStream) ReadBits(numBits int) int64 {
	if numBits > bs.bitLength {
		numBits = bs.bitLength
	}

	var result int64
	for i := 0; i < numBits; i++ {
		if bs.str[i] {
			result |= (1 << (numBits - 1 - i))
		}

	}
	bs.str = bs.str[numBits:]
	bs.bitLength -= numBits

	return result
}
