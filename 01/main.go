package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := aoc_utils.ReadInput("01/example.txt")
	data := aoc_utils.ReadInput("01/data.txt")

	fmt.Printf("Solution 1: %d\n", solve1(input))
	fmt.Printf("Solution 2: %d\n", solve2(input))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", solve2(data))
}

func solve1(d []string) int {
	leftList, rightList := make([]int, 0), make([]int, 0)
	for _, line := range d {
		split := strings.Split(line, "   ")
		l, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, l)
		rightList = append(rightList, r)
	}
	slices.Sort(leftList)
	slices.Sort(rightList)

	result := 0
	for i := 0; i < len(leftList); i++ {
		result += int(math.Abs(float64(leftList[i] - rightList[i])))
	}
	return result
}

func solve2(d []string) int {
	leftList, amountMap := make([]int, 0), make(map[int]int)
	for _, line := range d {
		split := strings.Split(line, "   ")
		l, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, l)
		amountMap[r]++
	}

	result := 0
	for _, val := range leftList {
		result += amountMap[val] * val
	}

	return result
}
