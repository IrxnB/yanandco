package bitoperations

func GetNbit(number, n int) int {
	return (number >> n) & 1
}

func SetNBit(number, n, value int) int {
	mask := 1 << n
	if value == 1 {
		return number | mask
	} else {
		return number &^ mask
	}
}
