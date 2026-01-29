package main

import (
	"advent2025/services"
	"flag"
	"fmt"
)

func main() {
	day := flag.Int("day", 0, "day to print")
	flag.Parse()

	fmt.Println("Hello there, advent of code 2025")

	if *day == 0 || *day == 1 {
		fmt.Println("Day one")
		resultDayOnePartOne := services.DayOnePartOne()
		fmt.Println(resultDayOnePartOne)
		resultDayOnePartTwo := services.DayOnePartTwo()
		fmt.Println(resultDayOnePartTwo)
	}

	if *day == 0 || *day == 2 {
		fmt.Println("Day two")
		resultDayTwoPartOne := services.DayTwoPartOne()
		fmt.Println(resultDayTwoPartOne)
		resultDayTwoPartTwo := services.DayTwoPartTwo()
		fmt.Println(resultDayTwoPartTwo)
	}

	if *day == 0 || *day == 3 {
		fmt.Println("Day three")
		resultDayThreePartOne := services.DayThreePartOne()
		fmt.Println(resultDayThreePartOne)
		resultDayThreePartTwo := services.DayThreePartTwo()
		fmt.Println(resultDayThreePartTwo)
	}

	if *day == 0 || *day == 4 {
		fmt.Println("Day four")
		resultDayFourPartOne := services.DayFourPartOne()
		fmt.Println(resultDayFourPartOne)
		resultDayFourPartTwo := services.DayFourPartTwo()
		fmt.Println(resultDayFourPartTwo)
	}

	if *day == 0 || *day == 5 {
		fmt.Println("Day five")
		resultDayFivePartOne := services.DayFivePartOne()
		fmt.Println(resultDayFivePartOne)
		resultDayFivePartTwo := services.DayFivePartTwo()
		fmt.Println(resultDayFivePartTwo)
	}
}
