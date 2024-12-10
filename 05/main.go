package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := aoc_utils.Example()
	data := aoc_utils.Data()

	separators := aoc_utils.GetSeparators(input)
	ruleMap := createRuleMap(input[:separators[0]])
	result, wrongUpdates := solve1(ruleMap, input[separators[0]+1:])
	fmt.Printf("Solution 1: %d\n", result)
	fmt.Printf("Solution 2: %d\n", solve2(ruleMap, wrongUpdates))

	separators = aoc_utils.GetSeparators(data)
	ruleMap = createRuleMap(data[:separators[0]])
	result, wrongUpdates = solve1(ruleMap, data[separators[0]+1:])
	fmt.Printf("Solution 1: %d\n", result)
	fmt.Printf("Solution 2: %d\n", solve2(ruleMap, wrongUpdates))
}

func solve1(ruleMap map[string][]string, updates []string) (int, [][]string) {
	correctUpdates, wrongUpdates := make([][]string, 0), make([][]string, 0)
	for _, update := range updates {
		split := strings.Split(update, ",")
		broken := false
		for x, pageNumber := range split {
			r := ruleMap[pageNumber]
			if r != nil {
				for _, prevPages := range split[:x] {
					if slices.Contains(r, prevPages) {
						broken = true
						break
					}
				}
			}
			if broken {
				break
			}
		}
		if !broken {
			correctUpdates = append(correctUpdates, split)
		} else {
			wrongUpdates = append(wrongUpdates, split)
		}
	}

	result := 0
	for _, update := range correctUpdates {
		m, err := strconv.Atoi(update[len(update)/2])
		if err != nil {
			panic(err)
		}
		result += m
	}
	return result, wrongUpdates
}

func solve2(ruleMap map[string][]string, updates [][]string) int {
	for _, update := range updates {
		for i := 0; i < len(update)-1; i++ {
			for j := 1; j < len(update); j++ {
				r := ruleMap[update[j]]
				if r != nil {
					for k, prevPage := range update[:j] {
						if slices.Contains(r, prevPage) {
							update[k], update[j] = update[j], update[k]
							break
						}
					}
				}
			}
		}
	}

	result := 0
	for _, update := range updates {
		m, err := strconv.Atoi(update[len(update)/2])
		if err != nil {
			panic(err)
		}
		result += m
	}
	return result
}

func createRuleMap(rules []string) map[string][]string {
	ruleMap := make(map[string][]string)
	for _, rule := range rules {
		split := strings.Split(rule, "|")
		if ruleMap[split[0]] == nil {
			ruleMap[split[0]] = []string{split[1]}
		} else {
			ruleMap[split[0]] = append(ruleMap[split[0]], split[1])
		}
	}
	return ruleMap
}
