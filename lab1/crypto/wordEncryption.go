package crypto

import "fmt"

func EncryptStringMonoAlphabet(data string, key string) (string, error) {
	tcData, err := EncodeString(data)

	if err != nil {
		return "", err
	}

	tcKey, err := EncodeString(key)
	if err != nil {
		return "", err
	}

	return DecodeToString(EncryptMono(tcData, tcKey)), nil
}

func DecryptStringMonoAlphabet(data string, key string) (string, error) {
	tcData, err := EncodeString(data)

	if err != nil {
		return "", err
	}

	tcKey, err := EncodeString(key)
	if err != nil {
		return "", err
	}

	return DecodeToString(DecryptMono(tcData, tcKey)), nil
}
func EncryptStringPolyAlphabet(data string, key string, shift int) (string, error) {
	tcData, err := EncodeString(data)

	if err != nil {
		return "", err
	}

	tcKey, err := EncodeString(key)
	if err != nil {
		return "", err
	}

	return DecodeToString(EncryptWord(tcData, tcKey, shift)), nil
}

func DecryptStringPolyAlphabet(data string, key string, shift int) (string, error) {
	tcData, err := EncodeString(data)

	if err != nil {
		return "", err
	}

	tcKey, err := EncodeString(key)
	if err != nil {
		return "", err
	}

	return DecodeToString(DecryptWord(tcData, tcKey, shift)), nil
}

func EncryptWord(word []*TelegraphChar, key []*TelegraphChar, shift int) []*TelegraphChar {
	tcResult := make([]*TelegraphChar, 0, len(word))

	for pos, tc := range word {
		tcResult = append(tcResult, tc.Encrypt(key, pos+shift))
	}

	return tcResult
}

func DecryptWord(data []*TelegraphChar, key []*TelegraphChar, shift int) []*TelegraphChar {
	tcResult := make([]*TelegraphChar, 0, len(data))

	for pos, tc := range data {
		tcResult = append(tcResult, tc.Decrypt(key, pos+shift))
	}

	return tcResult
}

func EncryptMono(data []*TelegraphChar, key []*TelegraphChar) []*TelegraphChar {
	tcResult := make([]*TelegraphChar, 0, len(data))

	for _, tc := range data {
		tcResult = append(tcResult, tc.Encrypt(key, 0))
	}

	return tcResult
}

func DecryptMono(data []*TelegraphChar, key []*TelegraphChar) []*TelegraphChar {
	tcResult := make([]*TelegraphChar, 0, len(data))

	for _, tc := range data {
		tcResult = append(tcResult, tc.Decrypt(key, 0))
	}

	return tcResult
}

func EncryptBlock(data string, key string, shift int) (string, error) {

	shift = shift % len(key)

	tcData, err := EncodeString(data)
	if err != nil {
		return "", err
	}

	if len(tcData) != 4 {
		return "", fmt.Errorf("размер блока не равен 4: %v", len(data))
	}

	tcKey, err := EncodeString(key)
	if err != nil {
		return "", err
	}
	tcResult := make([]*TelegraphChar, 0, len(tcData))
	for pos, tc := range tcData {
		tcResult = append(tcResult, tc.Encrypt(tcKey, pos+shift))
	}

	index := 0

	for i := shift; i < len(tcKey); i++ {
		index += int(tcKey[i].GetByte())
	}

	index %= 4

	for i := 0; i < 3; i++ {
		cur := (i + index) % 4
		next := (cur + 1) % 4
		tcResult[next] = tcResult[next].Plus(tcResult[cur])
	}

	return DecodeToString(tcResult), nil
}

func DecryptBlock(data string, key string, shift int) (string, error) {

	shift = shift % len(key)

	tcData, err := EncodeString(data)
	if err != nil {
		return "", err
	}

	if len(tcData) != 4 {
		return "", fmt.Errorf("размер блока не равен 4: %v", len(data))
	}

	tcKey, err := EncodeString(key)
	if err != nil {
		return "", err
	}

	index := 0

	for i := shift; i < len(tcKey); i++ {
		index += int(tcKey[i].GetByte())
	}

	index %= 4

	for i := 3; i > 0; i-- {
		cur := (i + index) % 4
		prev := (cur + 3) % 4
		tcData[cur] = tcData[cur].Minus(tcData[prev])
	}

	tcResult := make([]*TelegraphChar, 0, len(tcData))
	for pos, tc := range tcData {
		tcResult = append(tcResult, tc.Decrypt(tcKey, pos+shift))
	}

	return DecodeToString(tcResult), nil
}
