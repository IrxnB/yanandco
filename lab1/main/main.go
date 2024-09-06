package main

import (
	"fmt"
	alph "lab1/telegraph_alphabet"
)

func main() {
	for i := 'а'; i < 'а'+32; i++ {
		a, _ := alph.GetCharacterValue(i)
		fmt.Println(a)
	}
	a, _ := alph.GetCharacterValue(' ')
	println(a)
}
