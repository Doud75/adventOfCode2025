package services

import (
	"advent2025/utils"
	"sort"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

func DayFivePartOne() int {
	contentAsString := utils.ReadFile("./files/dayFive.txt")
	contentAsArray := strings.Split(contentAsString, "\n\n")

	rangeId := utils.StringToArray(contentAsArray[0], "\n")
	rangeMin, rangeMax := utils.SplitArrayOfStringBySep(rangeId, "-")
	ids := utils.StringToArray(contentAsArray[1], "\n")

	freshSum := 0

	for _, v := range ids {
		idCheck := utils.StringToInt(v)
		for i := 0; i < len(rangeMin); i++ {
			if idCheck >= utils.StringToInt(rangeMin[i]) && idCheck <= utils.StringToInt(rangeMax[i]) {
				freshSum += 1
				break
			}
		}
	}

	return freshSum
}

func DayFivePartTwo() int {
	contentAsString := utils.ReadFile("./files/dayFive.txt")
	contentAsArray := strings.Split(contentAsString, "\n\n")

	rangeId := utils.StringToArray(contentAsArray[0], "\n")
	rangeMin, rangeMax := utils.SplitArrayOfStringBySep(rangeId, "-")

	var intervals []Interval
	for i := 0; i < len(rangeMin); i++ {
		start := utils.StringToInt(rangeMin[i])
		end := utils.StringToInt(rangeMax[i])
		intervals = append(intervals, Interval{start, end})
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var merged []Interval
	current := intervals[0]

	for i := 1; i < len(intervals); i++ {
		next := intervals[i]

		if next.Start <= current.End {
			if next.End > current.End {
				current.End = next.End
			}
		} else {
			merged = append(merged, current)
			current = next
		}
	}
	merged = append(merged, current)

	freshPossible := 0
	for _, v := range merged {
		freshPossible += v.End - v.Start + 1
	}

	return freshPossible
}
