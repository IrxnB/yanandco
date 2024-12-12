package oneway

import (
	"yanandco/lab1/crypto"
	"yanandco/lab2/sblockint"
)

type SBlockInt = sblockint.SBlockInt
type SBlock = crypto.SBlock
type TelegraphChar = crypto.TelegraphChar

func OneWayEncryptSBlock(data *SBlock, key string, iterations int) (*SBlock, error) {
	if key == "" {
		key = "amazing_key"
	}
	var result SBlock
	for i := 0; i < iterations; i++ {
		iterationResult := SBlock{Chars: make([]*TelegraphChar, len(data.Chars))}
		for i := range data.Chars {
			iterationResult.Chars[i] = &TelegraphChar{Char: data.Chars[i].Char}
		}
		iterationResult.Encrypt(key, i)
		for i := range iterationResult.Chars {
			iterationResult.Chars[i] = iterationResult.Chars[i].Plus(data.Chars[i])
		}
		result = iterationResult
		key = iterationResult.ToString()
	}
	return &result, nil
}

func OneWayEncryptSBlockInt(data *SBlockInt, key string, iterations int) (*SBlockInt, error) {
	s_block, err := data.ToSBlock()
	if err != nil {
		return nil, err
	}
	s_block_encrypted, err := OneWayEncryptSBlock(s_block, key, iterations)
	if err != nil {
		return nil, err
	}
	res, err := sblockint.NewSBlockIntFromSBlock(s_block_encrypted)
	if err != nil {
		return nil, err
	}
	return res, nil
}
