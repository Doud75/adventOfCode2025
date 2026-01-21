package services

import (
	"advent2025/utils"
)

func DayFourPartOne() int {
	contentAsString := utils.ReadFile("./files/dayFour.txt")
	contentAsArray := utils.StringToArray(contentAsString, "\n")
	contentAsArrayOfByte := utils.ArrayStringToArrayByte(contentAsArray)

	sum := 0

	for i, _ := range contentAsArrayOfByte {
		for j, _ := range contentAsArrayOfByte[i] {
			count := 0
			if contentAsArrayOfByte[i][j] == '.' {
				continue
			}
			if i != 0 {
				if j != 0 {
					if checkRoll(i-1, j-1, contentAsArrayOfByte) {
						count += 1
					}
				}
				if j != len(contentAsArrayOfByte[i])-1 {
					if checkRoll(i-1, j+1, contentAsArrayOfByte) {
						count += 1
					}
				}
				if checkRoll(i-1, j, contentAsArrayOfByte) {
					count += 1
				}
			}
			if i != len(contentAsArrayOfByte)-1 {
				if j != 0 {
					if checkRoll(i+1, j-1, contentAsArrayOfByte) {
						count += 1
					}
				}
				if j != len(contentAsArrayOfByte[i])-1 {
					if checkRoll(i+1, j+1, contentAsArrayOfByte) {
						count += 1
					}
				}
				if checkRoll(i+1, j, contentAsArrayOfByte) {
					count += 1
				}
			}
			if j != 0 {
				if checkRoll(i, j-1, contentAsArrayOfByte) {
					count += 1
				}
			}
			if j != len(contentAsArrayOfByte[i])-1 {
				if checkRoll(i, j+1, contentAsArrayOfByte) {
					count += 1
				}
			}
			if count < 4 {
				sum += 1
			}
		}
	}
	return sum
}

func DayFourPartTwo() int {
	contentAsString := utils.ReadFile("./files/dayFour.txt")
	contentAsArray := utils.StringToArray(contentAsString, "\n")
	contentAsArrayOfByte := utils.ArrayStringToArrayByte(contentAsArray)

	quit := false
	totalNumber := 0
	var toRemove [][]int

	for quit == false {
		sum := 0
		for i, _ := range contentAsArrayOfByte {
			for j, _ := range contentAsArrayOfByte[i] {
				count := 0
				if contentAsArrayOfByte[i][j] == '.' {
					continue
				}
				if i != 0 {
					if j != 0 {
						if checkRoll(i-1, j-1, contentAsArrayOfByte) {
							count += 1
						}
					}
					if j != len(contentAsArrayOfByte[i])-1 {
						if checkRoll(i-1, j+1, contentAsArrayOfByte) {
							count += 1
						}
					}
					if checkRoll(i-1, j, contentAsArrayOfByte) {
						count += 1
					}
				}
				if i != len(contentAsArrayOfByte)-1 {
					if j != 0 {
						if checkRoll(i+1, j-1, contentAsArrayOfByte) {
							count += 1
						}
					}
					if j != len(contentAsArrayOfByte[i])-1 {
						if checkRoll(i+1, j+1, contentAsArrayOfByte) {
							count += 1
						}
					}
					if checkRoll(i+1, j, contentAsArrayOfByte) {
						count += 1
					}
				}
				if j != 0 {
					if checkRoll(i, j-1, contentAsArrayOfByte) {
						count += 1
					}
				}
				if j != len(contentAsArrayOfByte[i])-1 {
					if checkRoll(i, j+1, contentAsArrayOfByte) {
						count += 1
					}
				}
				if count < 4 {
					sum++
					toRemove = append(toRemove, []int{i, j})
				}
			}
		}
		if sum == 0 {
			quit = true
		}
		totalNumber += sum
		for _, v := range toRemove {
			contentAsArrayOfByte[v[0]][v[1]] = '.'
		}
	}
	return totalNumber
}

func checkRoll(i int, j int, row [][]byte) bool {
	return row[i][j] == '@'
}
