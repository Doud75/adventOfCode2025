package services

import (
	"advent2025/utils"
	"math/big"
)

func DayThreePartOne() int {
	maxSumVoltage := 0
	contentAsString := utils.ReadFile("./files/dayThree.txt")
	contentAsArray := utils.StringToArray(contentAsString, "\n")
	for i := 0; i < len(contentAsArray); i++ {
		bank := utils.StringToArrayOfInt(contentAsArray[i])
		firstMax := utils.MaxIndex(bank, 0, len(bank)-1)
		secondMax := utils.MaxIndex(bank, firstMax+1, len(bank))

		firstIndex, secondIndex := utils.Sort2Digit(firstMax, secondMax)

		maxVoltage := utils.JoinTwoNumbers(bank[firstIndex], bank[secondIndex])

		maxSumVoltage += maxVoltage
	}
	return maxSumVoltage
}

func DayThreePartTwo() *big.Int {
	maxSumVoltage := big.NewInt(0)
	contentAsString := utils.ReadFile("./files/dayThree.txt")
	contentAsArray := utils.StringToArray(contentAsString, "\n")
	for i := 0; i < len(contentAsArray); i++ {
		bank := utils.StringToArrayOfInt(contentAsArray[i])
		start := 0
		remaning := 12
		var indexTemp []int

		for remaning > 0 {
			end := len(bank) - remaning + 1
			maxIndex := utils.MaxIndex(bank, start, end)

			indexTemp = append(indexTemp, bank[maxIndex])
			start = maxIndex + 1
			remaning--
		}

		value := utils.DigitsToBigInt(indexTemp)
		maxSumVoltage.Add(maxSumVoltage, value)
	}
	return maxSumVoltage
}
