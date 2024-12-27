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
	id          [9]*TelegraphChar
	message     *bitstream.BitStream
	senderMac   [8]*TelegraphChar
	receiverMac [8]*TelegraphChar
	sec         *Block
	iv          *Block
	key         *Block
	packageType [2]*TelegraphChar
	iterations  int
	roundKeys   [][]*TelegraphChar
	cmac        *Block
}

func NewSession(
	packageType string,
	senderMac string,
	recieverMac string,
	key string,
	iterations int,
) (Session, error) {
	if len(senderMac) != 8 {
		return Session{}, errors.New("Wrong sender mac length: must be 8")
	}
	if len(recieverMac) != 8 {
		return Session{}, errors.New("Wrong reciever mac length: must be 8")
	}

	sessionSenderMac := make([]*TelegraphChar, 8)
	sessionRecieverMac := make([]*TelegraphChar, 8)

	sessionKey, _ := blockencryption.NewBlockFromString(key)
	seeds := blockencryption.GenerateSeeds(sessionKey)
	lgenerator, _ := generators.LinearComposition(seeds, generators.AlternatingLSFR)
	generator := *lgenerator
	sessionRoundKeys := blockencryption.GenerateKeys(lgenerator, iterations+1)
	sessionId := make([]*TelegraphChar, 9)
	for i := 0; i < 9; i++ {
		newChars, _ := generator().ToSBlock()
		sessionId[i] = newChars.Chars[0]
	}

	sessionPackageType := make([]*TelegraphChar, 2)
	if packageType == "31" {
		sessionPackageType[0] = &TelegraphChar{Char: 3}
		sessionPackageType[1] = &TelegraphChar{Char: 1}
	} else if packageType == "32" {
		sessionPackageType[0] = &TelegraphChar{Char: 3}
		sessionPackageType[1] = &TelegraphChar{Char: 2}
	} else {
		return Session{}, errors.New("Wrong package type")
	}

	for i, c := range senderMac {
		tc, _ := crypto.NewTelegraphChar(c)
		sessionSenderMac[i] = tc
	}
	for i, c := range recieverMac {
		tc, _ := crypto.NewTelegraphChar(c)
		sessionRecieverMac[i] = tc
	}

	sec := append(sessionPackageType, sessionSenderMac...)
	sec = append(sec, sessionRecieverMac...)
	sec = append(sec, sessionPackageType...)
	sec = append(sec, sessionId...)
	constadd := make([]*TelegraphChar, 5)
	for i := 0; i < 5; i++ {
		constadd[i] = &TelegraphChar{Char: 0}
	}
	sec = append(sec, constadd...)
	sessionSec, _ := blockencryption.NewBlockFromTelegraphChars(sec)
	iv, _ := blockencryption.NewBlockFromString("                ") //init as 0
	session := Session{
		packageType: [2]*TelegraphChar(sessionPackageType),
		id:          [9]*TelegraphChar(sessionId),
		key:         sessionKey,
		sec:         sessionSec,
		senderMac:   [8]*TelegraphChar(sessionSenderMac),
		receiverMac: [8]*TelegraphChar(sessionRecieverMac),
		iv:          iv,
		roundKeys:   sessionRoundKeys,
		iterations:  iterations,
	}
	session.cmac = session.CFB(session.sec, iv)

	return session, nil
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
	s.EAXCFB(pack)

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

func (s Session) EAXCFB(pack Package) {
	assData := make([]*TelegraphChar, 16)
	for i := 0; i < 2; i++ {
		assData[i] = &TelegraphChar{Char: pack.packageType[i].Char}
	}
	for i := 2; i < 11; i++ {
		assData[i] = &TelegraphChar{Char: pack.sessionId[i].Char}
	}
	for i := 11; i < 16; i++ {
		assData[i] = &TelegraphChar{Char: pack.length[i].Char}
	}
	assDataBlock, _ := blockencryption.NewBlockFromTelegraphChars(assData)
	assDataBlocks := make([]Block, 2)
	assDataBlocks[0] = assDataBlock.Copy() // how can I do that even assDataBlock is a link not an object
	assDataBlocks[1] = s.sec.Copy()

	civ := s.CFB(assDataBlocks, s.iv)
	tmp := s.CFB(pack.data, civ)
	// mac := xor(xor(tmp, civ)civ)

}

func (s Session) EAXCFBinv() {

}
