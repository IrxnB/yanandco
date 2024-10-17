package generators

import (
	"yanandco/lab2/oneway"
	"yanandco/lab2/sblockint"
)

type SBlockInt = sblockint.SBlockInt

type Generator = func() *SBlockInt

func LSFR(seed *SBlockInt) *Generator {
	return nil
}

func alternatingLSFR(seed *SBlockInt) *Generator {
	seeds := make([]sblockint.SBlockInt, 3)
	seeds[0] = *seed
	for i := 1; i < 3; i++ {
		new_seed, err := oneway.OneWayEncryptSBlockInt(seed, "a", i)
		if err != nil {
			panic(err)
		}
		seeds[i] = *new_seed
	}
	generators := make([]Generator, 3)
	for i := 0; i < 3; i++ {
		generators[i] = *LSFR(&seeds[i])
	}

	f := func() *SBlockInt {
		bit_arrays := make([]SBlockInt, 3)
		for i := 0; i < 3; i++ {
			bit_arrays[i] = *generators[i]()
		}

		res := int32(0)
		for i := 19; i >= 0; i-- {
			control_bit := bit_arrays[0].n >> i
			if (control_bit) == 0 {
				res += bit_arrays[1].n >> i
				res <<= 1
			} else {
				res += bit_arrays[2].n >> i
				res <<= 1
			}
		}
		return &SBlockInt{n: res}
	}
	return &f
}
