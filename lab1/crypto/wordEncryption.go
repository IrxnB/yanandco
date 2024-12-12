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

	return ToString(EncryptMono(tcData, tcKey)), nil
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

	return ToString(DecryptMono(tcData, tcKey)), nil
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

	return ToString(EncryptWord(tcData, tcKey, shift)), nil
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

	return ToString(DecryptWord(tcData, tcKey, shift)), nil
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

func (block *SBlock) Encrypt(key string, shift int) error {

	shift = shift % len(key)

	tcKey, err := EncodeString(key)
	if err != nil {
		return err
	}

	for pos, tc := range block.Chars {
		block.Chars[pos] = tc.Encrypt(tcKey, pos+shift)
	}

	index := 0

	for i := shift; i < len(tcKey); i++ {
		index += int(tcKey[i].GetByte())
	}

	index %= len(block.Chars)

	for i := 0; i < len(block.Chars); i++ {
		cur := (i + index) % len(block.Chars)
		next := (cur + 1) % len(block.Chars)
		block.Chars[next] = block.Chars[next].Plus(block.Chars[cur])
	}

	return nil
}

func (block *SBlock) Decrypt(key string, shift int) error {

	shift = shift % len(key)

	tcKey, err := EncodeString(key)
	if err != nil {
		return err
	}

	index := 0

	for i := shift; i < len(tcKey); i++ {
		index += int(tcKey[i].GetByte())
	}

	index %= 4

	for i := len(block.Chars) - 1; i > 0; i-- {
		cur := (i + index) % len(block.Chars)
		prev := (cur + 3) % len(block.Chars)
		block.Chars[cur] = block.Chars[cur].Minus(block.Chars[prev])
	}

	for pos, tc := range block.Chars {
		block.Chars[pos] = tc.Decrypt(tcKey, pos+shift)
	}

	return nil
}
