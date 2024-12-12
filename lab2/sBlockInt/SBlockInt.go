package sblockint

import (
	"fmt"
	"yanandco/lab1/crypto"
)

type SBlock = crypto.SBlock
type SBlockInt struct {
	n int
}

func (b *SBlockInt) GetValue() int {
	return b.n
}

func NewSBlockIntFromInt(n int) (*SBlockInt, error) {
	if n >= 1<<20 {
		return nil, fmt.Errorf("больше 20 бит")
	}
	return &SBlockInt{n}, nil
}

// ? should move this function to sblock.go file as
// ? func (data *SBlock) ToSBlockInt() (*SBlockInt, error)
func NewSBlockIntFromSBlock(data *SBlock) (*SBlockInt, error) {
	n := int(0)
	for _, char := range data.Chars {
		n <<= 5
		n += int(char.GetByte())
	}
	return &SBlockInt{n: int(n)}, nil
}

func (data *SBlockInt) ToSBlock() (*SBlock, error) {
	//? How do I put declared sblock length instead of 4 in the loop?
	data = &SBlockInt{n: data.n}
	var chars = make([]*crypto.TelegraphChar, 4)
	for i := 3; i >= 0; i-- {
		chars[i] = &crypto.TelegraphChar{Char: byte(data.n & 0x1f)}
		data.n >>= 5
	}

	return &SBlock{Chars: chars}, nil
}
