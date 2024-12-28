package test

import (
	"testing"
	"yanandco/lab1/crypto"
	"yanandco/lab3/blockencryption"
	"yanandco/lab4/sessionprotocol"
)

func TestPadding(t *testing.T) {
	packageType := make([]*crypto.TelegraphChar, 2)
	packageType[0] = &crypto.TelegraphChar{Char: 31}
	senderMac := make([]*crypto.TelegraphChar, 8)
	recieverMac := make([]*crypto.TelegraphChar, 8)
	sessionId := make([]*crypto.TelegraphChar, 9)
	ivData := make([]*crypto.TelegraphChar, 16)
	data := make([]*crypto.TelegraphChar, 23)
	macData := make([]*crypto.TelegraphChar, 16)

	for i := 0; i < 16; i++ {
		ivData[i] = &crypto.TelegraphChar{Char: 0}
		macData[i] = &crypto.TelegraphChar{Char: 0}
	}

	iv, _ := blockencryption.NewBlockFromTelegraphChars(ivData)
	mac, _ := blockencryption.NewBlockFromTelegraphChars(macData)
	pack := sessionprotocol.NewPackage(
		[2]*crypto.TelegraphChar(packageType),
		[8]*crypto.TelegraphChar(senderMac),
		[8]*crypto.TelegraphChar(recieverMac),
		[9]*crypto.TelegraphChar(sessionId),
		iv,
		data,
		mac,
	)

	bits, _ := pack.PadToBits()
	unpadded := sessionprotocol.Unpad(bits)

	t.Log(unpadded.GetPackageType()[0].Char)
}
