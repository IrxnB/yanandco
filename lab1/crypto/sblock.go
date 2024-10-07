package crypto

import (
	"fmt"
)

type SBlock struct {
	chars []*TelegraphChar
}

func NewSBlock(chars []*TelegraphChar) (*SBlock, error) {
	if len(chars) != 4 {
		return nil, fmt.Errorf("wrong number of chars")
	}
	return &SBlock{chars}, nil
}
