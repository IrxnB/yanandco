package generators

import (
	"fmt"
	"yanandco/lab2/bitoperations"
	"yanandco/lab2/oneway"
	"yanandco/lab2/sblockint"
)

type SBlockInt = sblockint.SBlockInt

type Generator = func() *SBlockInt

type GeneratorFactory = func(seed *sblockint.SBlockInt) *Generator

func LSFR(seed *SBlockInt) *Generator {
	state := seed.GetValue()

	lsfr := func() *sblockint.SBlockInt {
		for i := 1; i < 20; i++ {
			bit := bitoperations.GetNbit(state, 19) ^
				bitoperations.GetNbit(state, 18) ^
				bitoperations.GetNbit(state, 15) ^
				bitoperations.GetNbit(state, 13)

			state = bitoperations.SetNBit(state, 19, 0)

			state = state << 1

			state = bitoperations.SetNBit(state, 0, bit)
		}

		res, _ := sblockint.NewSBlockIntFromInt(state)
		return res
	}

	return &lsfr
}

func AlternatingLSFR(seed *SBlockInt) *Generator {
	seeds := make([]sblockint.SBlockInt, 3)
	seeds[0] = *seed
	for i := 1; i < 3; i++ {
		new_seed, err := oneway.OneWayEncryptSBlockInt(seed, "вдйсржхзцпчубъеякгтмшэлноф", i)
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

		res := 0
		for i := 19; i >= 0; i-- {
			control_bit := bitoperations.GetNbit(bit_arrays[0].GetValue(), i)
			if (control_bit) == 0 {
				res = bitoperations.SetNBit(res, i, bitoperations.GetNbit(bit_arrays[1].GetValue(), i))
			} else {
				res = bitoperations.SetNBit(res, i, bitoperations.GetNbit(bit_arrays[2].GetValue(), i))
			}
		}
		result, err := sblockint.NewSBlockIntFromInt(res)
		if err != nil {
			fmt.Println(res)
		}
		return result
	}
	return &f
}

func LinearComposition(seeds []*SBlockInt, factory GeneratorFactory) (*Generator, error) {
	if len(seeds) < 4 {
		return nil, fmt.Errorf("нужно 4 seed")
	}

	gen1 := *factory(seeds[0])
	gen2 := *factory(seeds[1])
	gen3 := *factory(seeds[2])
	gen4 := *factory(seeds[3])

	generator := func() *SBlockInt {
		res := gen1().GetValue() ^ gen2().GetValue() ^ gen3().GetValue() ^ gen4().GetValue()
		res <<= 12
		res >>= 12
		sblock, err := sblockint.NewSBlockIntFromInt(res)
		if err != nil {
			fmt.Println(err)
		}
		return sblock
	}
	return &generator, nil
}
