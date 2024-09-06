package telegraph_alphabet

import (
	"fmt"
	"math"
)

func GetCharacterValue(c rune) (byte, error) {
	if c == ' ' {
		return 0, nil
	}
	byte_k := c - 1071
	if byte_k > 32 {
		return 0, fmt.Errorf("out of range character: %v", c)
	}
	if byte_k == 29 { // 'ь' if input is 'ъ'
		return 27, nil
	}
	if byte_k > 29 {
		return byte(byte_k) - 1, nil
	}
	return byte(byte_k), nil
}

func Shift(c byte, shift byte) (byte, error) {
	if c == 0 {
		return 0, nil
	}
	c += shift
	c %= 31
	if c == 0 { // since we count from 1
		return 31, nil
	}
	return c, nil
}
