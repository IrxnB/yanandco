package test

import (
	"sync"
	"testing"
	"yanandco/lab2/generators"
	"yanandco/lab2/sblockint"
)

func TestBitShiftTrim(t *testing.T) {
	i := uint64(18446744073709551615)
	t.Log(i)

	t.Log(i << 44 >> 44)
}
func TestLFSR(t *testing.T) {
	seed, _ := sblockint.NewSBlockIntFromInt(1)
	gen := *generators.LSFR(seed)

	first := gen().GetValue()
	count := 0
	for {
		i := gen().GetValue()
		count++
		if i == first {
			break
		}
	}
	t.Logf("первое %v итераций %v", first, count)
}

func TestAlteratingLFSR(t *testing.T) {
	seed, _ := sblockint.NewSBlockIntFromInt(1)
	gen := *generators.AlternatingLSFR(seed)

	first := gen().GetValue()
	count := 0
	for {
		i := gen().GetValue()
		count++
		if i == first {
			break
		}
	}
	t.Logf("первое %v итераций %v", first, count)
}

func TestManyAlteratingLFSR(t *testing.T) {
	var wg sync.WaitGroup
	for i := 1000; i < 1100; i++ {
		wg.Add(1)

		go func(val int) {
			seed, _ := sblockint.NewSBlockIntFromInt(val)
			gen := *generators.AlternatingLSFR(seed)
			first := gen().GetValue()
			count := 0
			for {
				generated := gen().GetValue()
				count++
				if generated == first {
					break
				}
			}
			t.Logf("первое %v итераций %v", first, count)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestLinearComposition(t *testing.T) {
	seeds := make([]*generators.SBlockInt, 0, 4)

	seed1, _ := sblockint.NewSBlockIntFromInt(1)
	seed2, _ := sblockint.NewSBlockIntFromInt(2)
	seed3, _ := sblockint.NewSBlockIntFromInt(3)
	seed4, _ := sblockint.NewSBlockIntFromInt(4)

	seeds = append(seeds, seed1)
	seeds = append(seeds, seed2)
	seeds = append(seeds, seed3)
	seeds = append(seeds, seed4)

	genRef, _ := generators.LinearComposition(seeds, generators.AlternatingLSFR)

	gen := *genRef

	first := gen().GetValue()
	count := 0
	for {
		generated := gen().GetValue()
		count++
		if generated == first {
			break
		}
	}
	t.Logf("первое %v итераций %v", first, count)
}

func TestDispersionEquality(t *testing.T) {
	columnFactor := (1<<20-1)/10 + 1
	for seed_value := 1; seed_value < 5; seed_value++ {
		columns := make([]int, 10)
		seeds := make([]*generators.SBlockInt, 0, 4)

		seed1, _ := sblockint.NewSBlockIntFromInt(1 * seed_value)
		seed2, _ := sblockint.NewSBlockIntFromInt(2 * seed_value)
		seed3, _ := sblockint.NewSBlockIntFromInt(3 * seed_value)
		seed4, _ := sblockint.NewSBlockIntFromInt(4 * seed_value)

		seeds = append(seeds, seed1)
		seeds = append(seeds, seed2)
		seeds = append(seeds, seed3)
		seeds = append(seeds, seed4)
		genRef, _ := generators.LinearComposition(seeds, generators.AlternatingLSFR)
		gen := *genRef

		for i := 0; i < 1_000; i++ {
			repetition := gen()
			colNumber := repetition.GetValue() / columnFactor
			if colNumber >= 10 {
				t.Logf("")
			}
			columns[colNumber] += 1
		}
		for i, v := range columns {
			t.Logf("%v: %v", i*columnFactor, v)
		}
		t.Logf("---------------------------------------------------")
	}
}
