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

	result := make([]*crypto.TelegraphChar, len(data))

	for pos, val := range right {
		result[pos], _ = crypto.NewTelegraphChar(val.ToRune())
	}

	for pos, val := range left {
		result[pos+len(right)], _ = crypto.NewTelegraphChar(val.ToRune())
	}

	return result
}

func round(data []*crypto.TelegraphChar, key []*crypto.TelegraphChar, shift int) []*crypto.TelegraphChar {
	l0 := data[0:4]
	r0 := data[4:8]

	r1 := make([]*crypto.TelegraphChar, 4)

	for pos, val := range l0 {
		r1[pos], _ = crypto.NewTelegraphChar(val.ToRune())
	}

	l1sblock, _ := crypto.NewSBlockFromTC(l0)
	keyStr := crypto.ToString(key)
	l1sblock.Encrypt(keyStr, shift)
	l1 := l1sblock.Chars
	for i, v := range r0 {
		l1[i] = l1[i].Plus(v)
	}
	res := make([]*crypto.TelegraphChar, 8)

	for i := 0; i < 4; i++ {
		res[i], _ = crypto.NewTelegraphChar(l1[i].ToRune())
		res[i+4], _ = crypto.NewTelegraphChar(r1[i].ToRune())
	}

	return Skitala(res)
}

func (b *Block) Encrypt(key *Block, iterations int) error {
	seeds := make([]*sblockint.SBlockInt, 4)
	for i := 0; i < 4; i++ {
		seeds[i], _ = sblockint.NewSBlockIntFromSBlock(&crypto.SBlock{Chars: key.data[i*4 : i*4+4]})
	}

	generator, _ := generators.LinearComposition(seeds, generators.AlternatingLSFR)
	for i := 0; i < 4; i++ {
		_, _ = (*generator)().ToSBlock()
		for j := 0; j < 4; j++ {
			//b.data[i*4+j] = b.data[i*4+j].Xor(cur_key.Chars[j])
		}
	}

	for i := 0; i < iterations; i++ {
		curKey := make([]*crypto.TelegraphChar, 16)
		for i := 0; i < 4; i++ {
			cur_key, _ := (*generator)().ToSBlock()
			for j := 0; j < 4; j++ {
				curKey[i*4+j] = cur_key.Chars[j]
			}
		}
		fmt.Println(crypto.ToString(curKey))

		left := b.data[:8]
		right := b.data[8:]

		nextleft := round(left, curKey, i*4)
		for i, v := range nextleft {
			nextleft[i] = v.Xor(right[i])
		}

		nextright := left

		b.data = append(nextleft, nextright...)
	}
	fmt.Println("-------------")

	return nil
}

func (b *Block) Decrypt(key *Block, iterations int) error {
	seeds := make([]*sblockint.SBlockInt, 4)
	for i := 0; i < 4; i++ {
		seeds[i], _ = sblockint.NewSBlockIntFromSBlock(&crypto.SBlock{Chars: key.data[i*4 : i*4+4]})
	}
	generator, _ := generators.LinearComposition(seeds, generators.AlternatingLSFR)
	keys := make([]*Block, iterations+1)
	for i := 0; i < iterations+1; i++ {
		keys[i] = &Block{data: make([]*crypto.TelegraphChar, 16)}
		for j := 0; j < 4; j++ {
			sblock, _ := (*generator)().ToSBlock()
			for z := 0; z < 4; z++ {
				keys[i].data[j*4+z] = sblock.Chars[z]
			}
		}
	}

	for i := iterations - 1; i >= 0; i-- {
		left := b.data[:8]
		right := b.data[8:]
		fmt.Println(crypto.ToString(keys[i+1].data))

		prevright := round(right, keys[i+1].data, i*4)
		for i, v := range prevright {
			prevright[i] = v.Xor(left[i])
		}

		b.data = append(right, prevright...)
	}

	// for i, v := range b.data {
	// b.data[i] = v.Xor(keys[0].data[i])
	// }
	return nil
}
