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

	res := *data
	res.Chars = make([]*TelegraphChar, len(data.Chars))
	for i := range data.Chars {
		res.Chars[i] = &TelegraphChar{Char: data.Chars[i].Char}
	}

	for i := 0; i < iterations; i++ {

		buf := res
		buf.Chars = make([]*TelegraphChar, len(res.Chars))
		for i := range res.Chars {
			buf.Chars[i] = &TelegraphChar{Char: res.Chars[i].Char}
		}

		if err := res.Encrypt(key, i); err != nil {
			return nil, err
		}

		for j := 0; j < len(res.Chars); j++ {
			buf.Chars[j] = res.Chars[j].Plus(buf.Chars[j])
		}
		key = res.ToString()
	}
	return &res, nil
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
