package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToArray(str string, sep string) []string {
	return strings.Split(str, sep)
}

func ArrayStringToArrayByte(arr []string) [][]byte {
	grid := make([][]byte, len(arr))
	for i, line := range arr {
		grid[i] = []byte(line)
	}
	return grid
}

func ArrayStringToArrayInt(arr []string) []int {
	var result []int

	for _, line := range arr {
		n, err := strconv.Atoi(line)
		if err == nil {
			result = append(result, n)
		}
	}
	return result
}

func StringToInt(str string) int {
	intConverted, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return intConverted
}

func SplitArrayOfStringBySep(strArr []string, sep string) ([]string, []string) {
	var arrayLeft []string
	var arrayRight []string

	for i := 0; i < len(strArr); i++ {
		a := strings.Split(strArr[i], sep)
		arrayLeft = append(arrayLeft, a[0])
		arrayRight = append(arrayRight, a[1])
	}

	return arrayLeft, arrayRight
}

func SplitArrayOfStringByIndex(strArr []string, ind int) ([]string, []string) {
	var arrayLeft []string
	var arrayRight []string

	for i := 0; i < len(strArr); i++ {
		arrayLeft = append(arrayLeft, string(strArr[i][ind]))
		arrayRight = append(arrayRight, strArr[i][ind+1:])
	}

	return arrayLeft, arrayRight
}

func StringToArrayOfInt(str string) []int {
	var intArray []int

	for i := 0; i < len(str); i++ {
		intArray = append(intArray, StringToInt(string(str[i])))
	}

	return intArray
}

func JoinTwoNumbers(i int, j int) int {
	return i*10 + j
}
