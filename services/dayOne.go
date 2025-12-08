package services

import (
	"advent2025/utils"
)

func DayOnePartOne() int {
	dialNumber := 50
	count := 0
	contentAsString := utils.ReadFile("./files/dayOne.txt")
	contentAsArray := utils.StringToArray(contentAsString, "\n")
	direction, row := utils.SplitArrayOfStringByIndex(contentAsArray, 0)

	for i := 0; i < len(row); i++ {
		rowDigit := utils.StringToInt(row[i])
		if direction[i] == "R" {
			dialNumber = (dialNumber + rowDigit) % 100
		} else if direction[i] == "L" {
			dialNumber = (dialNumber - rowDigit + 100) % 100
		}

		if dialNumber == 0 {
			count++
		}
	}

	return count
}

func DayOnePartTwo() int {
	dialNumber := 50
	count := 0
	contentAsString := utils.ReadFile("./files/dayOne.txt")
	contentAsArray := utils.StringToArray(contentAsString, "\n")
	direction, row := utils.SplitArrayOfStringByIndex(contentAsArray, 0)

	for i := 0; i < len(row); i++ {
		rowDigit := utils.StringToInt(row[i])
		if direction[i] == "R" {
			count += (dialNumber + rowDigit) / 100
			dialNumber = (dialNumber + rowDigit) % 100
		} else if direction[i] == "L" {
			distToZero := dialNumber
			if distToZero == 0 {
				distToZero = 100
			}
			if rowDigit >= distToZero {
				count++
				remaining := rowDigit - distToZero
				count += remaining / 100
			}

			dialNumber = ((dialNumber-rowDigit)%100 + 100) % 100
		}
	}

	return count
}
