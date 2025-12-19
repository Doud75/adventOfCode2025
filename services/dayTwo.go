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
			NumberOfDigitPerBlock := numberOfDigit / 2

			power := utils.Power(NumberOfDigitPerBlock)

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
	contentAsString := utils.ReadFile("./files/dayTwoTest.txt")
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
			NumberOfDigitPerBlock := numberOfDigit / 2

			power := utils.Power(NumberOfDigitPerBlock)

			left := j / power
			right := j % power

			if left == right {
				invalidSumIds += j
			}
		}
	}
	return invalidSumIds
}
