package utils

func NumberOfDigits(N int) int {
	digits := 0
	for N > 0 {
		N /= 10
		digits++
	}
	return digits
}

func IsPair(N int) bool {
	return N%2 == 0
}

func Power(N int) int {
	powerInt := 1
	for i := 0; i < N; i++ {
		powerInt *= 10
	}
	return powerInt
}
