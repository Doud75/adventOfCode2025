package services

import (
	"advent2025/utils"
	"fmt"
)

func DayTwoPartOne() int {
	invalidSumIds := 0
	contentAsString := utils.ReadFile("./files/dayTwoTest.txt")
	contentAsArray := utils.StringToArray(contentAsString, ",")
	firstId, lastId := utils.SplitArrayOfStringBySep(contentAsArray, "-")
	for i := 0; i < len(firstId); i++ {
		fmt.Println(firstId[i])
		fmt.Println(lastId[i])
	}
	return invalidSumIds
}
