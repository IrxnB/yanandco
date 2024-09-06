package telegraph_alphabet

import (
	"fmt"
)

func EncodeToAlphabetByte(c rune) (byte, error) {
	if c == ' ' {
		return 0, nil
	}
	byte_k := c - 1071
	if byte_k > 32 {
		return 0, fmt.Errorf("character out of range ['а'-'я'] or ' ': %v", c)
	}
	if byte_k == 29 { // 'ь' if input is 'ъ'
		return 27, nil
	}
	if byte_k > 29 {
		return byte(byte_k) - 1, nil
	}
	return byte(byte_k), nil
}

func DecodeAlphabetByte(c byte) (rune, error) {
	if c > 31 {
		return 0, fmt.Errorf("character out of range [0 - 31] or ' ': %v", c)
	}
	if c == 0 {
		return ' ', nil
	}
	if c >= 29 {
		return rune(int(c) + 1072), nil
	}
	return rune(int(c) + 1071), nil
}

func Shift(c byte, shift byte) (byte, error) {
	if shift > 30 {
		return 0, fmt.Errorf("shift out of range [0; 30]: %v", shift)
	}
	if c == 0 {
		return 0, nil
	}
	c += shift
	c %= 31
	if c == 0 {
		return 31, nil
	}
	return c, nil
}

func EncodeString(s string) ([]byte, error) {
	enc_s := make([]byte, len(s))
	for i, c := range s {
		enc_c, err := EncodeToAlphabetByte(c)
		if err != nil {
			return nil, err
		}
		enc_s[i] = enc_c
	}
	return enc_s, nil
}

func DecodeString(s []byte) (string, error) {
	dec_s := make([]rune, len(s))
	for i, c := range s {
		dec_c, err := DecodeAlphabetByte(c)
		if err != nil {
			return "", err
		}
		dec_s[i] = dec_c
	}
	return string(dec_s), nil
}

func ShiftString(s string, shift byte) (string, error) {
	encodedString, err := EncodeString(s)
	if err != nil {
		return "", err
	}

	for i := range encodedString {
		shiftedByte, err := Shift(encodedString[i], shift)
		if err != nil {
			return "", err
		}
		encodedString[i] = shiftedByte
	}

	decodedString, err := DecodeString(encodedString)
	if err != nil {
		return "", err
	}

	return decodedString, nil

}
