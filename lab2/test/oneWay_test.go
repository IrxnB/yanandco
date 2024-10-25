package test

import (
	"reflect"
	"testing"
	"yanandco/lab1/crypto"
	"yanandco/lab2/oneway"
)

func TestOneWayEncrypt(t *testing.T) {
	s_start, err := crypto.NewSBlockFromString("вуду")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	s_end, err := oneway.OneWayEncryptSBlock(s_start, "", 4)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Before: %v, After 4 iterations: %v", s_start.ToString(), s_end.ToString())
	s_end, err = oneway.OneWayEncryptSBlock(s_end, "", 4)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Before: %v, After 8 iterations: %v", s_start.ToString(), s_end.ToString())
	if reflect.DeepEqual(s_end.Chars, s_start.Chars) {
		t.Error("Encryption failed")
		t.FailNow()
	}
}
