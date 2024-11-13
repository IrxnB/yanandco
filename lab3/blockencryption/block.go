package blockencryption

import (
	"fmt"
	"yanandco/lab1/crypto"
	"yanandco/lab2/generators"
	"yanandco/lab2/sblockint"
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

func NewBlock(seed *sblockint.SBlockInt) *Block {
	generator := *generators.AlternatingLSFR(seed)
	data := make([]*crypto.TelegraphChar, 16)
	for i := 0; i < 4; i++ {
		sblock, _ := generator().ToSBlock()
		for j := 0; j < 4; j++ {
			data[i*4+j] = sblock.Chars[j]
		}
	}
	return &Block{data: data}
}

// Это и есть P-блок по факту
func Skitala(data []*crypto.TelegraphChar) []*crypto.TelegraphChar {
	if len(data) < 4 || len(data)%2 != 0 {
		panic(fmt.Errorf("количество символов должно быть четным и не менее 4"))
	}
	result := make([]*crypto.TelegraphChar, len(data))

	right := data[:len(data)/2]
	left := data[len(data)/2:]

	for i := 0; i < len(data); i++ {
		if (i+1)%4 < 2 {
			result[i] = &crypto.TelegraphChar{Char: right[i/2].GetByte()}
		} else {
			result[i] = &crypto.TelegraphChar{Char: left[i/2].GetByte()}
		}
	}

	return result
}

// Обратный P-блок
func Antiskitala(data []*crypto.TelegraphChar) []*crypto.TelegraphChar {
	if len(data) < 4 || len(data)%2 != 0 {
		panic(fmt.Errorf("количество символов должно быть четным и не менее 4"))
	}

	right := make([]*crypto.TelegraphChar, len(data)/2)
	left := make([]*crypto.TelegraphChar, len(data)/2)

	for i := 0; i < len(data); i++ {
		if (i+1)%4 < 2 {
			right[i/2] = &crypto.TelegraphChar{Char: data[i].GetByte()}
		} else {
			left[i/2] = &crypto.TelegraphChar{Char: data[i].GetByte()}
		}
	}

	result := append(right, left...)
	return result
}

func round([]*crypto.TelegraphChar) []*crypto.TelegraphChar {
	return nil
}

// func (b *Block) encrypt(key *Block, iterations int) error {

// 	for i := 0; i < iterations; i++ {
// 		round
// }
// }
