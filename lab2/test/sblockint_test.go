package test

import (
	"math"
	"reflect"
	"testing"
	"yanandco/lab1/crypto"
	"yanandco/lab2/sblockint"
)

func TestSBlockIntFromSBlock(t *testing.T) {
	t.Setenv("GOTEST_TIMEOUT_SCALE", "100")
	for i := 1; i < int(math.Pow(31, 4)); i++ {
		chars := make([]*crypto.TelegraphChar, 4)
		for j := 0; j < 4; j++ {
			chars[j] = &crypto.TelegraphChar{Char: byte(i / int((math.Pow(31, float64(j)))) % 31)}
		}

		s_start := &crypto.SBlock{Chars: chars}

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

		t.Logf("Converted: %v, Start: %v, End: %v", sint, s_start.ToString(), s_end.ToString())
		if !reflect.DeepEqual(s_end.Chars, s_start.Chars) {
			t.Error("Conversion failed")
			t.FailNow()
		}
	}
}
