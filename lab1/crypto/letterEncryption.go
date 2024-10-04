package crypto

func (tc *TelegraphChar) Encrypt(key []*TelegraphChar, index int) *TelegraphChar {
	index = index % len(key)
	return tc.Plus(key[index])
}

func (tc *TelegraphChar) Decrypt(key []*TelegraphChar, index int) *TelegraphChar {
	index = index % len(key)
	return tc.Minus(key[index])
}
