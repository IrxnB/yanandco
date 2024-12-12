package sessionprotocol

import (
	"yanandco/lab1/crypto"
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
}

func (Session) SendMessage(message string) {
	return
}

func (Session) RecieveMessage() (string, error) {
	// перехватить ошибки если сообщение было повреждено
	return "", nil
}

func (Session) CFB() {
	// реализовывать сразу с CIV
	// Вынес сюда, потому что у Package нет SEC, но он есть непосредственно у получателя и отправителя
	return
}

func (Session) CFBinv() {
	return
}
