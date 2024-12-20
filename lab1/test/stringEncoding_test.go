package test

import (
	"strconv"
	"testing"
	"yanandco/lab1/crypto"
)

func TestEncoding(t *testing.T) {
	str := "какаято строкаь"
	encoded, err := crypto.EncodeString(str)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}

	decoded := crypto.ToString(encoded)

	decodedRuneArr := []rune(decoded)
	t.Logf("Изначальная строка: %v,\n после раскодирования %v", strconv.Quote(str), strconv.Quote(decoded))
	for pos, char := range []rune(str) {
		decodedChar := decodedRuneArr[pos]
		if decodedChar != char && char != 'ь' {
			t.FailNow()
		}
	}
}
