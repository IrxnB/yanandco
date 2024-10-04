package test

import (
	"testing"
	"yanandco/lab1/crypto"
)

func TestBlockEncryption(t *testing.T) {
	data := "блоа"
	key := "вдйсржхзцпчубъеякгтмшэлноф"
	shift := 2
	encrypted, err := crypto.EncryptBlock(data, key, shift)
	if err != nil {
		t.FailNow()
	}
	decrypted, err := crypto.DecryptBlock(encrypted, key, shift)
	if err != nil {
		t.FailNow()
	}

	t.Logf("%v, %v, %v", data, encrypted, decrypted)
}
