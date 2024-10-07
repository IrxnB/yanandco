package generators

import (
	"yanandco/lab2/sBlockInt"
)

type SBlockInt = sblockint.SBlockInt

type Generator = func() *SBlockInt

func LSFR(seed *SBlockInt) *Generator {
	return nil
}
