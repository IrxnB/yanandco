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

func TestAntiskitala(t *testing.T) {
	data := "адебвжзг"
	encoded, err := crypto.EncodeString(data)
	if err != nil {
		t.Fatalf("encode failed: %v", err)
		t.FailNow()
	}
	antiskitaled := blockencryption.Antiskitala(encoded)

	antiskitaled_string := crypto.ToString(antiskitaled)

	t.Logf("Start: %v, Encoded: %v", data, antiskitaled_string)
	if antiskitaled_string != "абвгдежз" {
		t.Fail()
	}
}

func TestEncrypt(t *testing.T) {
	block_data := "абвгдежзийклмноп"
	key_data := "йцукенгшзхъфывау"

	block, err := blockencryption.NewBlockFromString(block_data)
	if err != nil {
		t.Fatalf("block creation failed: %v", err)
		t.FailNow()
	}
	block_key, err := blockencryption.NewBlockFromString(key_data)
	if err != nil {
		t.Fatalf("key creation failed: %v", err)
		t.FailNow()
	}
	err = block.Encrypt(block_key, 2)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
		t.FailNow()
	}

	encrypted_string := block.ToString()
	t.Logf("Start: %v, Encrypted: %v", block_data, encrypted_string)
}
