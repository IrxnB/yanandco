package telegraphchar

import (
	"fmt"
)

type TelegraphChar struct {
	char byte
}

func FromRune(r rune) (*TelegraphChar, error) {
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
	if tc.char == 0 {
		return ' '
	}

	if tc.char >= 29 {
		return rune(tc.char) + 1072
	}

	return rune(tc.char) + 1071
}

func (tc *TelegraphChar) GetByte() byte {
	return tc.char
}

func (tc *TelegraphChar) Plus(another *TelegraphChar) *TelegraphChar {
	return &TelegraphChar{(tc.char + another.char) % 32}
}
