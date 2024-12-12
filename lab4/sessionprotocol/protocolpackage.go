package sessionprotocol

import (
	"yanandco/lab3/blockencryption"
)

type Block = blockencryption.Block

type Package struct {
	packageType *[2]TelegraphChar
	senderMac   *[8]TelegraphChar
	recieverMac *[8]TelegraphChar
	data        *[]TelegraphChar //так как блоки могут быть нецелыми. Можно еще со string попробовать
	iv          *[]TelegraphChar
	mac         Block
}

func (Package) toBin() []byte {
	// очень сомневаюсь в необходимости работать с битами. Как будто бы если и работать с ними, то в
	// SBlockInt переводить нужно все данные
	return make([]byte, 0)
}

func FromBin() Package {
	return Package{}
}

func (Package) packageLength() int {
	return 0
}

func (Package) padData() {
	return
}

func (Package) checkPadding() bool {
	return false
}

func (Package) unpadData() {
	return
}
