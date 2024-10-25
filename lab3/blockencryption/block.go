package block

import (
	"fmt"
	"yanandco/lab1/crypto"
)

type Block struct {
	data []*crypto.TelegraphChar
}

func NewBlockFromTelegraphChars(data []*crypto.TelegraphChar) (*Block, error) {
	if len(data) != 16 {
		return nil, fmt.Errorf("wrong number of chars")
	}
	return &Block{data: data}, nil
}

func NewBlockFromString(data string) (*Block, error) {
	chars, err := crypto.EncodeString(data)
	if err != nil {
		return nil, err
	}
	block, err := NewBlockFromTelegraphChars(chars)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (block *Block) ToString() string {
	return crypto.ToString(block.data)
}

func NewBlock(seed string) *Block {
	return nil
}

func skitala() []*crypto.TelegraphChar {
	return nil
}

func round() []*crypto.TelegraphChar {
	return nil, nil
}

func (b *Block) encrypt(key *Block, iterations int) error {
	return nil
}
