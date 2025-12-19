package main

import (
	"advent2025/services"
	"fmt"
)

func main() {
	fmt.Println("Hello there, advent of code 2025")
	fmt.Println("Day one")
	resultDayOnePartOne := services.DayOnePartOne()
	fmt.Println(resultDayOnePartOne)
	resultDayOnePartTwo := services.DayOnePartTwo()
	fmt.Println(resultDayOnePartTwo)

	fmt.Println("Day two")
	resultDayTwoPartOne := services.DayTwoPartOne()
	fmt.Println(resultDayTwoPartOne)
	resultDayTwoPartTwo := services.DayTwoPartTwo()
	fmt.Println(resultDayTwoPartTwo)

	fmt.Println("Day three")
	resultDayThreePartOne := services.DayThreePartOne()
	fmt.Println(resultDayThreePartOne)
	resultDayThreePartTwo := services.DayThreePartTwo()
	fmt.Println(resultDayThreePartTwo)
}
