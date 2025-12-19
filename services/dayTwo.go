package services

import (
	"advent2025/utils"
)

func DayTwoPartOne() int {
	invalidSumIds := 0
	contentAsString := utils.ReadFile("./files/dayTwo.txt")
	contentAsArray := utils.StringToArray(contentAsString, ",")
	firstId, lastId := utils.SplitArrayOfStringBySep(contentAsArray, "-")
	for i := 0; i < len(firstId); i++ {
		lastIntId := utils.StringToInt(lastId[i])
		firstIntId := utils.StringToInt(firstId[i])
		for j := firstIntId; j <= lastIntId; j++ {
			numberOfDigit := utils.NumberOfDigits(j)
			if !utils.IsPair(numberOfDigit) {
				continue
			}
			numberOfDigitPerBlock := numberOfDigit / 2

			power := utils.Power(numberOfDigitPerBlock, 10)

			left := j / power
			right := j % power

			if left == right {
				invalidSumIds += j
			}
		}
	}
	return invalidSumIds
}

func DayTwoPartTwo() int {
	invalidSumIds := 0
	contentAsString := utils.ReadFile("./files/dayTwo.txt")
	contentAsArray := utils.StringToArray(contentAsString, ",")
	firstId, lastId := utils.SplitArrayOfStringBySep(contentAsArray, "-")
	for i := 0; i < len(firstId); i++ {
		lastIntId := utils.StringToInt(lastId[i])
		firstIntId := utils.StringToInt(firstId[i])
		for j := firstIntId; j <= lastIntId; j++ {
			numberOfDigit := utils.NumberOfDigits(j)

			for numberOfDigitPerBlock := 1; numberOfDigitPerBlock <= numberOfDigit/2; numberOfDigitPerBlock++ {
				if numberOfDigit%numberOfDigitPerBlock != 0 {
					continue
				}

				r := numberOfDigit / numberOfDigitPerBlock
				if r < 2 {
					continue
				}
				power := utils.Power(numberOfDigitPerBlock, 10)
				block := j / utils.Power(r-1, power)

				rebuilt := 0
				for k := 0; k < r; k++ {
					rebuilt = rebuilt*power + block
				}

				if rebuilt == j {
					invalidSumIds += j
					break
				}
			}
		}
	}
	return invalidSumIds
}
