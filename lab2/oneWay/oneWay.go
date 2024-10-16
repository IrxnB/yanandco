package oneway

import (
	"yanandco/lab1/crypto"
)

type SBlock = crypto.SBlock
type TelegraphChar = crypto.TelegraphChar

func OneWayEncrypt(data *SBlock, key string, iterations int) (*SBlock, error) {
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
		buf.Chars = make([]*TelegraphChar, len(data.Chars))
		for i := range data.Chars {
			buf.Chars[i] = &TelegraphChar{Char: data.Chars[i].Char}
		}

		if err := data.Encrypt(key, i); err != nil {
			return nil, err
		}
		for j := 0; j < len(data.Chars); j++ {
			buf.Chars[j] = data.Chars[j].Plus(buf.Chars[j])
		}
		key = data.ToString()
	}
	return &res, nil
}
