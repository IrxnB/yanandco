package test

import (
	"reflect"
	"testing"
	"yanandco/lab1/crypto"
	"yanandco/lab2/sblockint"
)

func TestSBlockIntFromSBlock(t *testing.T) {
	s_start, err := crypto.NewSBlockFromString("вуду")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	sint, err := sblockint.NewSBlockIntFromSBlock(s_start)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	s_end, err := sint.ToSBlock()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if !reflect.DeepEqual(s_end.Chars, s_start.Chars) {
		t.Error("Conversion failed")
		t.FailNow()
	}
	t.Logf("Converted: %v, Start: %v, End: %v", sint, s_start.ToString(), s_end.ToString())
}
