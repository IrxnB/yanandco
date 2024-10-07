package sblockint

import (
	"yanandco/lab1/crypto"
)

type SBlock = crypto.SBlock

type SBlockInt struct {
	n int
}

func ConvertBlock(data *SBlock) (*SBlockInt, error) {
	return &SBlockInt{n: 0}, nil
}

func ConvertInt(data *SBlockInt) (*SBlock, error) {
	return nil, nil
}
