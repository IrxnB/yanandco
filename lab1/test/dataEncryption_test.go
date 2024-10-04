package test

import (
	"lab1/crypto"
	"strconv"
	"testing"
)

func TestEncodingMonoAlphabet(t *testing.T) {
	testStr := "какаято строка"
	keyStr := "ключ строка"

	encoded, err := crypto.EncryptStringMonoAlphabet(testStr, keyStr)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	decoded, err := crypto.DecryptStringMonoAlphabet(encoded, keyStr)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	runeArrDecoded := []rune(decoded)
	for pos, char := range []rune(testStr) {
		if runeArrDecoded[pos] != char && char != 'Ь' {
			t.FailNow()
		}
	}

	t.Logf(
		"\nИзначальная %v, ключ %v, закодированная %v, раскодированная %v",
		strconv.Quote(testStr),
		strconv.Quote(keyStr),
		strconv.Quote(encoded),
		strconv.Quote(decoded))
}

func TestEncodingPolyAlphabet(t *testing.T) {
	testStr := "какаято строка"
	keyStr := "ключ строка"

	encoded, err := crypto.EncryptStringPolyAlphabet(testStr, keyStr, 0)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	decoded, err := crypto.DecryptStringPolyAlphabet(encoded, keyStr, 0)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	runeArrDecoded := []rune(decoded)
	for pos, char := range []rune(testStr) {
		if runeArrDecoded[pos] != char && char != 'Ь' {
			t.FailNow()
		}
	}
	t.Logf(
		"\nИзначальная %v, ключ %v, закодированная %v, раскодированная %v",
		strconv.Quote(testStr),
		strconv.Quote(keyStr),
		strconv.Quote(encoded),
		strconv.Quote(decoded))
}
