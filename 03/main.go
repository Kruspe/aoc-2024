package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"regexp"
	"strconv"
)

func main() {
	input := aoc_utils.ReadInput("03/example.txt")
	input2 := aoc_utils.ReadInput("03/example2.txt")
	data := aoc_utils.ReadInput("03/data.txt")

	fmt.Printf("Solution 1: %d\n", solve1(input))
	fmt.Printf("Solution 2: %d\n", solve2(input2))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", solve2(data))
}

func solve1(d []string) int {
	mulRe, err := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	if err != nil {
		panic(err)
	}
	numRe, err := regexp.Compile("[0-9]+")
	if err != nil {
		panic(err)
	}
	result := 0
	for _, line := range d {
		res := mulRe.FindAllString(line, -1)
		for _, r := range res {
			numbers := numRe.FindAllString(r, -1)
			n1, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
			n2, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}
			result += n1 * n2
		}
	}

	return result
}

func solve2(d []string) int {
	doRe, err := regexp.Compile("do\\(\\)")
	if err != nil {
		panic(err)
	}
	dontRe, err := regexp.Compile("don't\\(\\)")
	if err != nil {
		panic(err)
	}
	mulRe, err := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	if err != nil {
		panic(err)
	}
	numRe, err := regexp.Compile("[0-9]+")
	if err != nil {
		panic(err)
	}

	result, state := 0, "do"
	for _, line := range d {
		currentStopIndex, currentStartIndex, startStopIndex := 0, 0, make([][]int, 0)
		doIndex := doRe.FindAllStringIndex(line, -1)
		dontIndex := dontRe.FindAllStringIndex(line, -1)

		var startDo *int
		for {
			if state == "dont" {
				if len(doIndex) == currentStartIndex {
					break
				}
				if doIndex[currentStartIndex][0] > dontIndex[currentStopIndex][1] {
					state = "do"
					startDo = &doIndex[currentStartIndex][1]
				} else {
					currentStartIndex++
				}
			} else {
				if len(dontIndex) == currentStopIndex {
					break
				}
				if startDo == nil {
					state = "dont"
					startStopIndex = append(startStopIndex, []int{0, dontIndex[currentStopIndex][0]})
					continue
				}
				if dontIndex[currentStopIndex][0] > doIndex[currentStartIndex][1] {
					state = "dont"
					startStopIndex = append(startStopIndex, []int{doIndex[currentStartIndex][1], dontIndex[currentStopIndex][0]})
					startDo = nil
				} else {
					currentStopIndex++
				}
			}
		}
		if startDo != nil {
			startStopIndex = append(startStopIndex, []int{*startDo, len(line)})
		}

		sanitizedLine := ""
		for _, index := range startStopIndex {
			sanitizedLine += line[index[0]:index[1]]
		}

		res := mulRe.FindAllString(sanitizedLine, -1)
		for _, r := range res {
			numbers := numRe.FindAllString(r, -1)
			n1, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
			n2, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}
			result += n1 * n2
		}
	}
	return result
}
