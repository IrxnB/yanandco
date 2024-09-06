package telegraph_alphabet

import (
	"fmt"
)

func EncodeToAlphabetByte(c rune) (byte, error) {
	if c == ' ' {
		return 0, nil
	}
	k_byte := c - 1071
	if k_byte > 32 {
		return 0, fmt.Errorf("character out of range ['а'-'я'] or ' ': %v", c)
	}
	if k_byte == 29 { // 'ь' if input is 'ъ'
		return 27, nil
	}
	if k_byte > 29 {
		return byte(k_byte) - 1, nil
	}
	return byte(k_byte), nil
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
	s_rune := []rune(s)
	enc_s := make([]byte, len(s_rune))
	for i := 0; i < len(s_rune); i++ {
		enc_c, err := EncodeToAlphabetByte(s_rune[i])
		if err != nil {
			return nil, err
		}
		enc_s[i] = enc_c
	}
	return enc_s, nil
}

func DecodeString(s_byte []byte) (string, error) {
	dec_s := make([]rune, len(s_byte))
	for i, c := range s_byte {
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
