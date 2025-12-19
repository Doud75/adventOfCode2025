package utils

import "math/big"

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

func Power(N int, a int) int {
	powerInt := 1
	for i := 0; i < N; i++ {
		powerInt *= a
	}
	return powerInt
}

func MaxIndex(arr []int, startIndex int, end int) int {
	maxInt := arr[startIndex]
	index := startIndex
	for i := startIndex; i < end; i++ {
		if arr[i] > maxInt {
			maxInt = arr[i]
			index = i
		}
	}
	return index
}

func Sort2Digit(i int, j int) (int, int) {
	if i < j {
		return i, j
	}
	return j, i
}

func DigitsToBigInt(digits []int) *big.Int {
	var result big.Int
	for _, d := range digits {
		result.Mul(&result, big.NewInt(10))
		result.Add(&result, big.NewInt(int64(d)))
	}
	return &result
}
