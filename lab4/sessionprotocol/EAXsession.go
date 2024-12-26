package sessionprotocol

import (
	"yanandco/lab1/crypto"
	"yanandco/lab2/generators"
	"yanandco/lab3/blockencryption"
)

type TelegraphChar = crypto.TelegraphChar

type Session struct {
	id          TelegraphChar
	message     []byte
	senderMac   TelegraphChar
	receiverMac TelegraphChar
	sec         string
	iv          TelegraphChar
	mac         Block
	key         Block
	iterations  int
	roundKeys   [][]*crypto.TelegraphChar
}

func CreateSession(key Block, iterations int) Session {
	seeds := blockencryption.GenerateSeeds(&key)
	generator, _ := generators.LinearComposition(seeds, generators.AlternatingLSFR)
	roundKeys := blockencryption.GenerateKeys(generator, iterations+1)

	//Посчитать SEC и тд
	return Session{roundKeys: roundKeys}
}

func (Session) SendMessage(message string) {
	// create package
	// encode

	return
}

func (Session) RecieveMessage() (string, error) {
	// перехватить ошибки если сообщение было повреждено
	return "", nil
}

func (s Session) CFB(data []Block, iv Block) []Block {
	blocks := len(data)
	result := make([]Block, blocks)
	prev := iv.Copy()
	for i := 0; i < blocks; i++ {
		prev.EncryptPregen(s.roundKeys, s.iterations)
		result[i] = data[i].Xor(prev)
		prev = result[i].Copy()
	}
	return result
}

func (s Session) CFBinv(data []Block, iv Block) []Block {
	blocks := len(data)
	result := make([]Block, blocks)
	prev := iv.Copy()
	for i := blocks - 1; i >= 0; i++ {
		prev.EncryptPregen(s.roundKeys, s.iterations)
		result[i] = data[i].Xor(prev)
		prev = data[i].Copy()
	}
	return result
}
