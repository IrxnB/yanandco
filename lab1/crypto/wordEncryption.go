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

	for pos, tc := range block.chars {
		block.chars[pos] = tc.Encrypt(tcKey, pos+shift)
	}

	index := 0

	for i := shift; i < len(tcKey); i++ {
		index += int(tcKey[i].GetByte())
	}

	index %= len(block.chars)

	for i := 0; i < len(block.chars); i++ {
		cur := (i + index) % len(block.chars)
		next := (cur + 1) % len(block.chars)
		block.chars[next] = block.chars[next].Plus(block.chars[cur])
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

	for i := len(block.chars) - 1; i > 0; i-- {
		cur := (i + index) % len(block.chars)
		prev := (cur + 3) % len(block.chars)
		block.chars[cur] = block.chars[cur].Minus(block.chars[prev])
	}

	for pos, tc := range block.chars {
		block.chars[pos] = tc.Decrypt(tcKey, pos+shift)
	}

	return nil
}
