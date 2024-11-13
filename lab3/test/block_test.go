package test

import (
	"testing"
	"yanandco/lab1/crypto"
	"yanandco/lab3/blockencryption"
)

func TestSkitala(t *testing.T) {
	data := "абвгдежз"
	encoded, err := crypto.EncodeString(data)
	if err != nil {
		t.Fatalf("encode failed: %v", err)
		t.FailNow()
	}
	skitaled := blockencryption.Skitala(encoded)

	skitaled_string := crypto.ToString(skitaled)

	t.Logf("Start: %v, Encoded: %v", data, skitaled_string)
	if skitaled_string != "адебвжзг" {
		t.Fail()
	}
}
