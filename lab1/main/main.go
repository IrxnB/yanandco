package main

import (
	"fmt"
	alph "lab1/telegraph_alphabet"
)

func main() {
	s := "а привет всем я крутой программист кстати но настраивать чтение из строки было впадлу"

	shift := byte(5)
	shiftback := byte(31 - shift)

	s, err := alph.ShiftString(s, shift)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
	s, err = alph.ShiftString(s, shiftback)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
