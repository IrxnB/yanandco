package telegraph_alphabet

import "fmt"

func GetCharacterValue(k rune) (byte, error) {
	if k == ' ' {
		return 0, nil
	}
	byte_k := k - 1071
	if byte_k > 32 {
		return 0, fmt.Errorf("out of range character: %v", k)
	}
	if byte_k == 29 { // 'ь' if input is 'ъ'
		return 27, nil
	}
	if byte_k > 29 {
		return byte(byte_k) - 1, nil
	} else {
		return byte(byte_k), nil
	}
}
