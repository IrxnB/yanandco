package test

import (
	"testing"
	"yanandco/lab1/crypto"
)

func TestBlockEncryption(t *testing.T) {
	data := "блоа"
	key := "вдйсржхзцпчубъеякгтмшэлноф"
	shift := 2

	block, err := crypto.NewSBlockFromString(data)
	if err != nil {
		t.FailNow()
	}

	err = block.Encrypt(key, shift)
	if err != nil {
		t.FailNow()
	}

	encrypted := block.ToString()

	err = block.Decrypt(key, shift)
	if err != nil {
		t.FailNow()
	}

	decrypted := block.ToString()

	t.Logf("%v, %v, %v", data, encrypted, decrypted)
}
