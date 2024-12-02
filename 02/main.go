package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := aoc_utils.ReadInput("02/example.txt")
	data := aoc_utils.ReadInput("02/data.txt")

	fmt.Printf("Solution 1: %d\n", solve1(input))
	fmt.Printf("Solution 2: %d\n", solve2(input))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", solve2(data))
}

func solve1(d []string) int {
	deltas := make([][]int, 0)
	for _, line := range d {
		deltas = append(deltas, createDeltas(line))
	}

	result := len(d)
	for _, delta := range deltas {
		increasing := delta[0] > 0
		for _, val := range delta {
			abs := int(math.Abs(float64(val)))
			if (increasing && val < 0) || (!increasing && val > 0) || abs < 1 || abs > 3 {
				result -= 1
				break
			}
		}
	}
	return result
}

func solve2(d []string) int {
	deltaMap := make(map[int][]int)
	for i, line := range d {
		deltaMap[i] = createDeltas(line)
	}

	result := len(d)
	recoveryMap := make(map[int][]int)
	for i, delta := range deltaMap {
		increasing := delta[0] > 0
		for j, val := range delta {
			abs := int(math.Abs(float64(val)))
			isUnsafe := (increasing && val < 0) || (!increasing && val > 0) || abs < 1 || abs > 3
			if isUnsafe {
				if recoveryMap[i] == nil {
					recoveryMap[i] = []int{j}
				} else {
					recoveryMap[i] = append(recoveryMap[i], j)
				}
			}

		}
	}

	recoveryDeltas := make(map[int][][]int)
	for index, recovery := range recoveryMap {
		moreDeltas := make([][]int, 0)
		levels := strings.Split(d[index], " ")
		for _, recoveryIndex := range recovery {
			l := append([]string{}, levels[:recoveryIndex]...)
			l = append(l, levels[recoveryIndex+1:]...)
			l2 := append([]string{}, levels[:recoveryIndex+1]...)
			l2 = append(l2, levels[recoveryIndex+2:]...)
			moreDeltas = append(moreDeltas, createDeltas(strings.Join(l, " ")))
			moreDeltas = append(moreDeltas, createDeltas(strings.Join(l2, " ")))
		}
		recoveryDeltas[index] = moreDeltas
	}

	for _, deltas := range recoveryDeltas {
		recovered := false
		for _, delta := range deltas {
			increasing := delta[0] > 0
			recoverable := true
			for _, val := range delta {
				abs := int(math.Abs(float64(val)))
				isUnsafe := (increasing && val < 0) || (!increasing && val > 0) || abs < 1 || abs > 3
				if isUnsafe {
					recoverable = false
					break
				}
			}
			if recoverable {
				recovered = true
				break
			}
		}
		if !recovered {
			result -= 1
		}
	}

	return result
}

func createDeltas(levelsLine string) []int {
	levels := strings.Split(levelsLine, " ")
	delta := make([]int, 0, len(levels)-1)
	for i := 0; i < len(levels)-1; i++ {
		l, err := strconv.Atoi(levels[i])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(levels[i+1])
		delta = append(delta, r-l)
	}
	return delta
}
