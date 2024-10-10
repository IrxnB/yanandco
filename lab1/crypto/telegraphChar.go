package crypto

import (
	"bytes"
	"fmt"
)

type TelegraphChar struct {
	Char byte
}

func NewTelegraphChar(r rune) (*TelegraphChar, error) {
	if r == ' ' {
		return &TelegraphChar{0}, nil
	}
	charRune := r - 1071
	if charRune > 32 {
		return nil, fmt.Errorf("rune out of range ['а'-'я'] or ' ': %v", r)
	}
	if charRune == 29 { // 'ь' if input is 'ъ'
		return &TelegraphChar{27}, nil
	}
	if charRune > 29 {
		return &TelegraphChar{byte(charRune) - 1}, nil
	}
	return &TelegraphChar{byte(charRune)}, nil
}

func (tc *TelegraphChar) ToRune() rune {
	if tc.Char == 0 {
		return ' '
	}

	if tc.Char >= 29 {
		return rune(tc.Char) + 1072
	}

	return rune(tc.Char) + 1071
}

func (tc *TelegraphChar) GetByte() byte {
	return tc.Char
}

func (tc *TelegraphChar) Plus(another *TelegraphChar) *TelegraphChar {
	return &TelegraphChar{(tc.Char + another.Char) % 32}
}

func (tc *TelegraphChar) Minus(another *TelegraphChar) *TelegraphChar {
	return &TelegraphChar{(tc.Char - another.Char + 32) % 32}
}

func EncodeString(str string) ([]*TelegraphChar, error) {
	arr := make([]*TelegraphChar, 0, len([]byte(str)))

	for _, char := range str {
		tc, err := NewTelegraphChar(char)
		if err != nil {
			return nil, err
		}
		arr = append(arr, tc)
	}

	return arr, nil
}

func ToString(arr []*TelegraphChar) string {
	var bytes bytes.Buffer

	for _, tc := range arr {
		bytes.WriteRune(tc.ToRune())
	}

	return bytes.String()
}
