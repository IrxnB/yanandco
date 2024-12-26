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

func (p Package) toBin() *bitstream.BitStream {
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

	return bs
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

func Unpad(bits *bitstream.BitStream) Package {
	_, padLength := getBlockCountPadLengthIfPadValid(bits, bits.Length()/80)

	bits.ReadBits(padLength)

	return binPackage
}

func getBlockCountPadLengthIfPadValid(bits *bitstream.BitStream, blockCount int) (int, int) {
	copy := bits.Copy()
	end := copy.ReadBits(3)
	if end != 0b001 {
		return 0, 0
	}
	n := copy.ReadBits(10)
	if n != blockCount {
		return 0, 0
	}
	l := copy.ReadBits(7)
	placeholder := l - 23
	for i := 0; i < placeholder; i++ {
		zero := copy.ReadBits(1)
		if zero != 0 {
			return 0, 0
		}
	}
	start := copy.ReadBits(3)
	if start != 0b100 {
		return 0, 0
	}
	return n, l
}

func (p Package) Pad() []*Block {
	bits := p.toBin()
	blockCount := bits.Length() / 80
	rem := bits.Length() % 80
	if rem != 0 {
		appendix, blockInc := createPadding(rem, blockCount)
		blockCount += blockInc
		bits.Append(appendix)
	} else if n, _ := getBlockCountPadLengthIfPadValid(bits, blockCount); n != 0 {
		appendix, blockInc := createPadding(80-23, blockCount)
		blockCount += blockInc
		bits.Append(appendix)
	}
	blocks := make([]*Block, blockCount)
	for i := 0; i < bits.Length()/80; i += 80 {
		blockData := make([]*TelegraphChar, 16)
		for j := 0; j < 16; j++ {
			blockData[j] = &TelegraphChar{Char: byte(bits.ReadBits(5))}
		}
		newblock, _ := blockencryption.NewBlockFromTelegraphChars(blockData)
		blocks = append(blocks, newblock)
	}

	return blocks
}

func createPadding(remainder int, initialBLocks int) (*bitstream.BitStream, int) {
	blockInc := 0
	placeholder := 0
	res := bitstream.NewBitStream()
	res.WriteBits(0b100, 3)
	if remainder < 23 && remainder > 0 {
		blockInc = 2
		placeholder = 80 + remainder
	}
	if remainder >= 23 {
		blockInc = 1
		placeholder = remainder
	}

	res.WriteBits(initialBLocks+blockInc, 10)
	res.WriteBits(placeholder+23, 7)

	for i := 0; i < placeholder; i++ {
		res.WriteBits(0b0, 1)
	}

	res.WriteBits(0b100, 3)
	return res, blockInc
}
