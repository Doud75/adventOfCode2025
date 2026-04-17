package services

import (
	"advent2025/utils"
	"strings"
)

func DaySixPartOne() int {
	contentAsString := utils.ReadFile("./files/daySix.txt")
	contentAsArray := utils.StringToArray(contentAsString, "\n")
	operator := strings.Fields(contentAsArray[len(contentAsArray)-1])
	result := 0
	var contentAsArrayOfArray [][]int
	for i := 0; i < len(contentAsArray)-1; i++ {
		fields := strings.Fields(contentAsArray[i])
		contentAsArrayOfArray = append(contentAsArrayOfArray, utils.ArrayStringToArrayInt(fields))
	}
	for i := 0; i < len(contentAsArrayOfArray[0]); i++ {
		tempResult := 0
		if operator[i] == "*" {
			tempResult = 1
		}
		for j := 0; j < len(contentAsArray)-1; j++ {
			if operator[i] == "*" {
				tempResult *= contentAsArrayOfArray[j][i]
			} else if operator[i] == "+" {
				tempResult += contentAsArrayOfArray[j][i]
			}
		}
		result += tempResult
	}
	return result
}
