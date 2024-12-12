module yanandco/lab4

go 1.21.5

replace yanandco/lab1 => ../lab1

replace yanandco/lab2 => ../lab2

replace yanandco/lab3 => ../lab3

require (
	yanandco/lab1 v0.0.0-00010101000000-000000000000
	yanandco/lab3 v0.0.0-00010101000000-000000000000
)

require yanandco/lab2 v0.0.0-00010101000000-000000000000 // indirect
