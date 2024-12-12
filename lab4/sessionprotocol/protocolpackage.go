package sessionprotocol

import (
	"yanandco/lab3/blockencryption"
)

type Block = blockencryption.Block

type Package struct {
	packageType [2]TelegraphChar
	senderMac   [8]TelegraphChar
	recieverMac [8]TelegraphChar
	sessionId   [9]TelegraphChar
	length      [5]TelegraphChar
	iv          *Block
	data        *[]TelegraphChar //так как блоки могут быть нецелыми. Можно еще со string попробовать
	mac         *Block
}

func (p Package) toBin() []byte {
	// Форма не бинарная, просто унифицируем все данные под byte. 3 бита каждого байта не используются
	// C таким же успехом (и по-моему логичнее) можно было хранить массив TelegraphChar
	binPackage := make([]byte, p.packageLength())
	binPackage[0] = p.packageType[0].GetByte()
	binPackage[1] = p.packageType[1].GetByte()
	offset := 2
	for i := 0; i < len(p.senderMac); i++ {
		binPackage[i+offset] = p.senderMac[i].GetByte()
	}
	offset += len(p.senderMac)

	for i := 0; i < len(p.recieverMac); i++ {
		binPackage[i+offset] = p.senderMac[i].GetByte()
	}
	offset += len(p.recieverMac)

	for i := 0; i < len(p.sessionId); i++ {
		binPackage[i+offset] = p.sessionId[i].GetByte()
	}
	offset += len(p.sessionId)

	for i := 0; i < len(p.length); i++ {
		binPackage[i+offset] = p.length[i].GetByte()
	}
	offset += len(p.length)

	for i := 0; i < len(p.iv.Data); i++ {
		binPackage[i+offset] = p.iv.Data[i].GetByte()
	}
	offset += len(p.iv.Data)

	for i := 0; i < len(*p.data); i++ {
		binPackage[i+offset] = (*p.data)[i].GetByte()
	}
	offset += len(*p.data)

	for i := 0; i < len(p.mac.Data); i++ {
		binPackage[i+offset] = p.mac.Data[i].GetByte()
	}
	return make([]byte, 0)
}

func FromBin(binPackage []byte) *Package {
	// TODO handle padding
	p := &Package{}
	packageType := [2]TelegraphChar{}
	packageType[0] = TelegraphChar{Char: binPackage[0]}
	packageType[1] = TelegraphChar{Char: binPackage[1]}
	p.packageType = packageType
	offset := 2

	senderMac := [8]TelegraphChar{}
	for i := 0; i < len(senderMac); i++ {
		senderMac[i] = TelegraphChar{Char: binPackage[i+offset]}
	}
	p.senderMac = senderMac
	offset += len(senderMac)

	recieverMac := [8]TelegraphChar{}
	for i := 0; i < len(recieverMac); i++ {
		recieverMac[i] = TelegraphChar{Char: binPackage[i+offset]}
	}
	p.recieverMac = recieverMac
	offset += len(recieverMac)

	sessionId := [9]TelegraphChar{}
	for i := 0; i < len(sessionId); i++ {
		sessionId[i] = TelegraphChar{Char: binPackage[i+offset]}
	}
	p.sessionId = sessionId
	offset += len(sessionId)

	lengthData := [5]TelegraphChar{}
	for i := 0; i < len(lengthData); i++ {
		lengthData[i] = TelegraphChar{Char: binPackage[i+offset]}
	}
	p.length = lengthData
	offset += len(lengthData)

	ivData := make([]*TelegraphChar, 16)
	for i := 0; i < len(ivData); i++ {
		ivData[i] = &TelegraphChar{Char: binPackage[i+offset]}
	}
	iv, _ := blockencryption.NewBlockFromTelegraphChars(ivData)
	p.iv = iv
	offset += len(ivData)

	data := make([]TelegraphChar, 0)
	dataLen := 0
	for i := 0; i < len(p.length); i++ {
		dataLen += int(p.length[i].GetByte())
	}
	for i := 0; i < dataLen; i++ {
		data = append(data, TelegraphChar{Char: binPackage[i+offset]})
	}
	p.data = &data
	offset += dataLen

	macData := make([]*TelegraphChar, 16)
	for i := 0; i < len(macData); i++ {
		macData[i] = &TelegraphChar{Char: binPackage[i+offset]}
	}
	mac, _ := blockencryption.NewBlockFromTelegraphChars(macData)
	p.mac = mac

	return p
}

func (p Package) packageLength() int {
	return len(p.packageType) + len(p.senderMac) + len(p.recieverMac) + len(*p.data) + 1 + len(p.mac.Data)
}

func (Package) padData() []byte {
	return
}

func (Package) checkPadding() bool {
	return false
}

func (Package) unpadData() {
	return
}
