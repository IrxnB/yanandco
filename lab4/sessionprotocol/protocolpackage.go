package sessionprotocol

import (
	"yanandco/lab1/crypto"
	"yanandco/lab3/blockencryption"
	"yanandco/lab4/bitstream"
)

type Block = blockencryption.Block

type Package struct {
	packageType [2]TelegraphChar
	senderMac   [8]TelegraphChar
	recieverMac [8]TelegraphChar
	sessionId   [9]TelegraphChar
	length      [5]TelegraphChar //можно в int переделать
	iv          *Block
	data        *[]TelegraphChar
	mac         *Block
}

func (p Package) toBin() bitstream.BitStream {
	bs := bitstream.NewBitStream()

	for _, tc := range p.packageType {
		bs.WriteTelegraphChar(tc)
	}
	for _, tc := range p.senderMac {
		bs.WriteTelegraphChar(tc)
	}
	for _, tc := range p.recieverMac {
		bs.WriteTelegraphChar(tc)
	}
	for _, tc := range p.sessionId {
		bs.WriteTelegraphChar(tc)
	}
	for _, tc := range p.length {
		bs.WriteTelegraphChar(tc)
	}
	for _, tc := range p.iv.Data {
		bs.WriteTelegraphChar(*tc)
	}
	for _, tc := range *p.data {
		bs.WriteTelegraphChar(tc)
	}
	for _, tc := range p.mac.Data {
		bs.WriteTelegraphChar(*tc)
	}

	return *bs
}

func FromBin(bs bitstream.BitStream) *Package {
	p := &Package{}

	for i := 0; i < 2; i++ {
		p.packageType[i].Char = byte(bs.ReadBits(5))
	}

	for i := 0; i < 8; i++ {
		p.senderMac[i].Char = byte(bs.ReadBits(5))
	}

	for i := 0; i < 8; i++ {
		p.recieverMac[i].Char = byte(bs.ReadBits(5))
	}

	for i := 0; i < 9; i++ {
		p.sessionId[i].Char = byte(bs.ReadBits(5))
	}

	for i := 0; i < 5; i++ {
		p.length[i].Char = byte(bs.ReadBits(5))
	}

	ivData := make([]*crypto.TelegraphChar, 16)
	for i := 0; i < 16; i++ {
		ivData[i] = &crypto.TelegraphChar{Char: byte(bs.ReadBits(5))}
	}
	p.iv, _ = blockencryption.NewBlockFromTelegraphChars(ivData)

	dataLen := p.dataLength()

	data := make([]crypto.TelegraphChar, dataLen)
	for i := 0; i < dataLen; i++ {
		data[i].Char = byte(bs.ReadBits(5))
	}
	p.data = &data

	macData := make([]*crypto.TelegraphChar, 16)
	for i := 0; i < 16; i++ {
		macData[i] = &crypto.TelegraphChar{Char: byte(bs.ReadBits(5))}
	}

	p.mac, _ = blockencryption.NewBlockFromTelegraphChars(macData)

	return p
}

// bit length of data
func (p Package) dataLength() int {
	dataLen := 0
	for _, char := range p.length {
		dataLen += int(char.Char)
	}
	dataLen *= 5
	return dataLen
}

// bit length of package
func (p Package) packageLength() int {
	total := 0
	total += len(p.packageType) * 5
	total += len(p.senderMac) * 5
	total += len(p.recieverMac) * 5
	total += len(p.sessionId) * 5
	total += len(p.length) * 5
	total += len(p.iv.Data) * 5
	total += p.dataLength()
	total += len(p.mac.Data) * 5
	return total
}

func (p Package) padMessage() bitstream.BitStream {
	l := p.packageLength()
	blocks := l / 80
	remainder := l % 80
	binPackage := p.toBin()

	padSize := 0
	if remainder == 0 {
		blocks += 1
		emptSize = 80
	} else if remainder <= 57 {
		block += 1
		padSize = 80 - remainder
	} else {
		blocks += 2
		padSize = 160 - remainder
	}

	return
}

func (Package) checkPadding() bool {
	return false
}

func (Package) unpadData() {
	return
}
