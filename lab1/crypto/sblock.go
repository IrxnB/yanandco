package crypto

import (
	"fmt"
)

type SBlock struct {
	Chars []*TelegraphChar
}

func NewSBlockFromTC(chars []*TelegraphChar) (*SBlock, error) {
	if len(chars) != 4 {
		return nil, fmt.Errorf("wrong number of chars")
	}
	return &SBlock{chars}, nil
}

func NewSBlockFromString(data string) (*SBlock, error) {
	chars, err := EncodeString(data)
	if err != nil {
		return nil, err
	}
	return NewSBlockFromTC(chars)
}

func (block *SBlock) ToString() string {
	return ToString(block.Chars)
}
