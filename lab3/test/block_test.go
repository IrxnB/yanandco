package test

import (
	"math/rand"
	"testing"
	"time"
	"yanandco/lab1/crypto"
	"yanandco/lab2/bitoperations"
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

// постоянный вход, меняем один бит
// произвольный вход и ключ
// меняем один бит в ключе и смотрим на изменение выхода
// меняем один бит
// построить гистограмму рассеивания

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
	err = block.Encrypt(block_key, 20)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
		t.FailNow()
	}
	encrypted_string := block.ToString()

	err = block.Decrypt(block_key, 20)
	if err != nil {
		t.Fatalf("dencryption failed: %v", err)
		t.FailNow()
	}

	decrypted_string := block.ToString()
	t.Logf("Start: %v, Encrypted: %v, Decrypted: %v", block_data, encrypted_string, decrypted_string)
}

func TestDispersion(t *testing.T) {
	block_data := "абвгдежзийклмноп"
	key_data := "абвгдежзийклмноп"

	data_changes_gistoram := make([]int, 10, 10)

	original_block, _ := blockencryption.NewBlockFromString(block_data)
	key, _ := blockencryption.NewBlockFromString(key_data)
	for i := 0; i < 80; i++ {
		encrypt_block, _ := blockencryption.NewBlockFromString(block_data)
		encrypt_block.Data[i/5] = &crypto.TelegraphChar{Char: byte(bitoperations.InverNBit(int(encrypt_block.Data[i/5].GetByte()), i%5))}
		encrypt_block.Encrypt(key, 6)

		changed := 0
		for j := 0; j < 16; j++ {
			difference := original_block.Data[j].Xor(encrypt_block.Data[j])
			for b := 0; b < 5; b++ {
				if (bitoperations.GetNbit(int(difference.Char), b)) == 1 {
					changed += 1
				}
			}
		}

		if changed/8 < 10 {
			data_changes_gistoram[changed/8] += 1
		}
	}
	t.Log(data_changes_gistoram)
}

func TestDispersionSequentialy(t *testing.T) {
	block_data := "абвгдежзийклмноп"
	key_data := "абвгдежзийклмноп"

	data_changes_gistoram := make([]int, 10, 10)

	prev_block, _ := blockencryption.NewBlockFromString(block_data)
	key, _ := blockencryption.NewBlockFromString(key_data)
	for i := 0; i < 1000; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		ind := r.Intn(80)
		next_block, _ := blockencryption.NewBlockFromTelegraphChars(make([]*crypto.TelegraphChar, len(prev_block.Data)))
		for j := range prev_block.Data {
			next_block.Data[j] = &crypto.TelegraphChar{Char: prev_block.Data[j].Char}
		}
		next_block.Data[ind/5] = &crypto.TelegraphChar{Char: byte(bitoperations.InverNBit(int(next_block.Data[ind/5].GetByte()), ind%5))}
		next_block.Encrypt(key, 6)

		changed := 0
		for j := 0; j < 16; j++ {
			difference := prev_block.Data[j].Xor(next_block.Data[j])
			for b := 0; b < 5; b++ {
				if (bitoperations.GetNbit(int(difference.Char), b)) == 1 {
					changed += 1
				}
			}
		}

		if changed/8 < 10 {
			data_changes_gistoram[changed/8] += 1
		}
		prev_block = next_block
	}
	t.Log(data_changes_gistoram)
}
