package crypto

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
func EncryptStringPolyAlphabet(data string, key string, shift rune) (string, error) {
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

func DecryptStringPolyAlphabet(data string, key string, shift rune) (string, error) {
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

func EncryptWord(word []*TelegraphChar, key []*TelegraphChar, shift rune) []*TelegraphChar {
	tcResult := make([]*TelegraphChar, 0, len(word))

	for pos, tc := range word {
		tcResult = append(tcResult, tc.Encrypt(key, pos+int(shift)))
	}

	return tcResult
}

func DecryptWord(word []*TelegraphChar, key []*TelegraphChar, shift rune) []*TelegraphChar {
	tcResult := make([]*TelegraphChar, 0, len(word))

	for pos, tc := range word {
		tcResult = append(tcResult, tc.Decrypt(key, pos+int(shift)))
	}

	return tcResult
}

func EncryptMono(word []*TelegraphChar, key []*TelegraphChar) []*TelegraphChar {
	tcResult := make([]*TelegraphChar, 0, len(word))

	for _, tc := range word {
		tcResult = append(tcResult, tc.Encrypt(key, 0))
	}

	return tcResult
}

func DecryptMono(word []*TelegraphChar, key []*TelegraphChar) []*TelegraphChar {
	tcResult := make([]*TelegraphChar, 0, len(word))

	for _, tc := range word {
		tcResult = append(tcResult, tc.Decrypt(key, 0))
	}

	return tcResult
}
