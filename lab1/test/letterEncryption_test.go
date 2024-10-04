package test

import (
	"bytes"
	"strconv"
	"testing"
	"yanandco/lab1/crypto"
)

func TestSingleLetterEncryption(t *testing.T) {
	var bytes bytes.Buffer

	for i := 'а'; i <= 'я'; i++ {
		bytes.WriteRune(i)
	}
	bytes.WriteRune(' ')

	key := bytes.String()
	keyTCArr, err := crypto.EncodeString(key)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for _, beforeEnc := range keyTCArr {
		for index := 0; index < len(keyTCArr); index++ {
			afterEnc := beforeEnc.Encrypt(keyTCArr, index)
			afterDec := afterEnc.Decrypt(keyTCArr, index)
			t.Logf(
				"До кодирования %v, после кодирования %v, после декодирования %v",
				strconv.QuoteRune(beforeEnc.ToRune()),
				strconv.QuoteRune(afterEnc.ToRune()),
				strconv.QuoteRune(afterDec.ToRune()))
			if afterDec.ToRune() != beforeEnc.ToRune() {
				t.FailNow()
			}
		}
	}
}
