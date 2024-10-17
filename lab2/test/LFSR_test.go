package test

import (
	"sync"
	"testing"
	"time"
	"yanandco/lab2/generators"
	"yanandco/lab2/oneway"
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

		time.Sleep(time.Millisecond * 500)
		go func(val int) {
			seed, _ := sblockint.NewSBlockIntFromInt(val)
			seed, _ = oneway.OneWayEncryptSBlockInt(seed, time.Now().String(), 10)
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
