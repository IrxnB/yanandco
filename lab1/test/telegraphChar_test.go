package test

import (
	"strconv"
	"testing"
	"yanandco/lab1/crypto"
)

func TestRuAlphabetEncoding(t *testing.T) {
	for i := 'а'; i <= 'я'; i++ {
		tc, err := crypto.NewTelegraphChar(i)

		if err != nil {
			t.Errorf("Конвертация буквы %v вызвала ошибку", string(i))
			t.Fail()
		}

		if i < 1100 && (tc.GetByte() != byte(i-1071)) {
			t.Errorf("Ожидаемое значение %v, получено значение %v", i-1071, tc.GetByte())
			t.Fail()
		} else if i == 1100 && (tc.GetByte() != 27) {
			t.Errorf("Ожидаемое значение %v, получено значение %v", 27, tc.GetByte())
			t.Fail()
		} else if i > 1100 && (tc.GetByte() != byte(i-1072)) {
			t.Errorf("Ожидаемое значение %v, получено значение %v", 1-1072, tc.GetByte())
			t.Fail()
		}
		if tc.ToRune() != i && i != 1100 {
			t.Errorf("Ожидаемое значение %v, получено значение %v", strconv.QuoteRune(i), strconv.QuoteRune(tc.ToRune()))
		}
		t.Logf("Буква: %v, TelegraphChar.ToRune(): %v", strconv.QuoteRune(i), strconv.QuoteRune(tc.ToRune()))
	}
}

func TestSpaceEncoding(t *testing.T) {
	tc, err := crypto.NewTelegraphChar(' ')
	if err != nil {
		t.Errorf("Конвертация буквы %v вызвала ошибку", strconv.QuoteRune(' '))
		t.Fail()
	}

	if tc.GetByte() != 0 {
		t.Errorf("Ожидаемое значение %v, получено значение %v", 0, tc.GetByte())
		t.Fail()
	}

	if tc.ToRune() != ' ' {
		t.Errorf("Ожидаемое значение %v, получено значение %v", strconv.QuoteRune(' '), strconv.QuoteRune(tc.ToRune()))
		t.Fail()
	}
	t.Logf("Буква: %v, TelegraphChar.ToRune(): %v", strconv.QuoteRune(' '), strconv.QuoteRune(tc.ToRune()))
}

func TestPlus(t *testing.T) {
	TCa, _ := crypto.NewTelegraphChar('а')

	TCb := TCa.Plus(TCa)

	t.Logf("Ожидаемое значение %v, получено значение %v", string('б'), string(TCb.ToRune()))
	if TCb.ToRune() != 'б' {
		t.Fail()
	}

	TCt, _ := crypto.NewTelegraphChar('т')

	TCe := TCt.Plus(TCt)

	t.Logf("Ожидаемое значение %v, получено значение %v", string('е'), string(TCe.ToRune()))
	if TCe.ToRune() != 'е' {
		t.Fail()
	}

}
