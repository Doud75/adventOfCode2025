package main

import (
    "fmt"
    "advent2025/services"
)

func main() {
    fmt.Println("Hello there, advent of code 2025")
    fmt.Println("Day one")
    resultDayOnePartOne := services.DayOnePartOne()
    fmt.Println(resultDayOnePartOne)
    resultDayOnePartTwo := services.DayOnePartTwo()
    fmt.Println(resultDayOnePartTwo)
}
