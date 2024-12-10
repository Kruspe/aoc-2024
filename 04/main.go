package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"regexp"
)

func main() {
	input := aoc_utils.Example()
	data := aoc_utils.Data()

	fmt.Printf("Solution 1: %d\n", solve1(input))
	fmt.Printf("Solution 2: %d\n", solve2(input))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", solve2(data))
}

func solve1(d []string) int {
	linesToSearch := make([]string, 0)
	for y, line := range d {
		linesToSearch = append(linesToSearch, line)
		backward := ""
		for _, c := range line {
			backward = string(c) + backward
		}
		linesToSearch = append(linesToSearch, backward)
		horizontal, backwardHorizontal := "", ""
		for i := 0; i < len(d[0]); i++ {
			horizontal += string(d[i][y])
			backwardHorizontal = string(d[i][y]) + backwardHorizontal
		}
		linesToSearch = append(linesToSearch, horizontal, backwardHorizontal)
	}
	for y := 0; y < len(d); y++ {
		diagonal, backwardDiagonal := "", ""
		for x := 0; x <= y; x++ {
			diagonal += string(d[y-x][x])
			backwardDiagonal = string(d[y-x][x]) + backwardDiagonal
		}
		linesToSearch = append(linesToSearch, diagonal, backwardDiagonal)
		diagonal, backwardDiagonal, xCounter := "", "", y
		for x := len(d) - 1; x >= len(d)-1-y; x-- {
			diagonal += string(d[x][xCounter])
			backwardDiagonal = string(d[x][xCounter]) + backwardDiagonal
			xCounter--
		}
		linesToSearch = append(linesToSearch, diagonal, backwardDiagonal)
	}
	for x := len(d) - 1; x >= 1; x-- {
		diagonal, backwardDiagonal := "", ""
		for y := len(d) - 1; y >= x; y-- {
			diagonal += string(d[y][x+(len(d)-1-y)])
			backwardDiagonal = string(d[y][x+(len(d)-1-y)]) + backwardDiagonal
		}
		linesToSearch = append(linesToSearch, diagonal, backwardDiagonal)

		diagonal, backwardDiagonal = "", ""
		for y := 0; y <= len(d)-1-x; y++ {
			diagonal += string(d[y][x+y])
			backwardDiagonal = string(d[y][x+y]) + backwardDiagonal
		}
		linesToSearch = append(linesToSearch, diagonal, backwardDiagonal)
	}

	re, err := regexp.Compile("XMAS")
	if err != nil {
		panic(err)
	}

	result := 0
	for _, line := range linesToSearch {
		matches := re.FindAllString(line, -1)
		result += len(matches)
	}

	return result
}

func solve2(d []string) int {
	result := 0
	for y, line := range d[1 : len(d)-1] {
		y = y + 1
		for x, c := range line[1 : len(line)-1] {
			x = x + 1
			if c == 'A' {
				topLeft, topRight, bottomRight, bottomLeft := d[y-1][x-1], d[y-1][x+1], d[y+1][x+1], d[y+1][x-1]
				if (topLeft == 'M' && bottomRight == 'S' || topLeft == 'S' && bottomRight == 'M') && (topRight == 'M' && bottomLeft == 'S' || topRight == 'S' && bottomLeft == 'M') {
					result += 1
				}
			}
		}
	}
	return result
}
