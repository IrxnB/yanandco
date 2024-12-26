package sessionprotocol

import (
	"errors"
	"yanandco/lab1/crypto"
	"yanandco/lab2/generators"
	"yanandco/lab3/blockencryption"
	"yanandco/lab4/bitstream"
)

type TelegraphChar = crypto.TelegraphChar

type Session struct {
	id          [9]TelegraphChar
	message     bitstream.BitStream
	senderMac   [8]TelegraphChar
	receiverMac [8]TelegraphChar
	sec         Block
	iv          Block
	key         Block
	packageType [2]TelegraphChar
	iterations  int
	roundKeys   [][]*crypto.TelegraphChar
}

func CreateSession(key Block, iterations int) Session {
	seeds := blockencryption.GenerateSeeds(&key)
	generator, _ := generators.LinearComposition(seeds, generators.AlternatingLSFR)
	

	//Посчитать SEC и тд
	return Session{roundKeys: roundKeys}
}

func NewSession(packageType string, senderMac string, recieverMac string, key string) (Session, error) {
	if len(senderMac) != 8 {
		return Session{}, errors.New("Wrong sender mac length: must be 8")
	}
	if len(recieverMac) != 8 {
		return Session{}, errors.New("Wrong reciever mac length: must be 8")
	}

	sessionSenderMac := make([]crypto.TelegraphChar, 8)
	sessionRecieverMac := make([]crypto.TelegraphChar, 8)

	sessionKey, _ := blockencryption.NewBlockFromString(key)
	seeds := blockencryption.GenerateSeeds(sessionKey)
  sessionRoundKeys := blockencryption.GenerateKeys(generator, iterations+1)
	lgenerator, _ := generators.LinearComposition(seeds, generators.AlternatingLSFR)
	generator := *lgenerator
	sessionId := make([]TelegraphChar, 9)
	for i := 0; i < 9; i++ {
		newChars, _ := generator().ToSBlock()
		sessionId[i] = *newChars.Chars[0]
	}

	sessionPackageType := make([]crypto.TelegraphChar, 2)
	if packageType == "31" {
		sessionPackageType[0] = TelegraphChar{Char: 3}
		sessionPackageType[1] = TelegraphChar{Char: 1}
	} else if packageType == "31" {
		sessionPackageType[0] = TelegraphChar{Char: 3}
		sessionPackageType[1] = TelegraphChar{Char: 1}
	} else {
		return Session{}, errors.New("Wrong package type")
	}

	for i, c := range senderMac {
		tc, _ := crypto.NewTelegraphChar(c)
		sessionSenderMac[i] = *tc
	}
	for i, c := range recieverMac {
		tc, _ := crypto.NewTelegraphChar(c)
		sessionRecieverMac[i] = *tc
	}

	// sec := append(sessionSenderMac, sessionRecieverMac...)
	// sec = append(sec, sessionPackageType...)
	// sec = append(sec, sessionId...)
	// constadd := make([]TelegraphChar, ?)

	// sessionSec = append(sessionSec, crypto.TelegraphChar{Char: 0})
	sessionSec, _ := blockencryption.NewBlockFromString("шестнадцатьсимво")
	iv, _ := blockencryption.NewBlockFromString("                ") //init as 0

	return Session{
		packageType: [2]crypto.TelegraphChar(sessionPackageType),
		id:          [9]TelegraphChar(sessionId),
		key:         *sessionKey,
		sec:         *sessionSec,
		senderMac:   [8]TelegraphChar(sessionSenderMac),
		receiverMac: [8]TelegraphChar(sessionRecieverMac),
		iv:          *iv,
    roundKeys:   sessionRoundKeys
	}, nil
}

func (session Session) SendMessage(message string) {
	rune_data := []rune(message)
	data := make([]TelegraphChar, len(rune_data))
	for i, r := range rune_data {
		tc, _ := crypto.NewTelegraphChar(r)
		data[i] = *tc
	}
	pack := NewPackage(
		session.packageType,
		session.receiverMac,
		session.senderMac,
		session.id,
		&session.iv,
		data,
		mac)
	session.iv.Data[0] = session.iv.Data[0].Plus(&TelegraphChar{Char: 1})

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
