package main

import (
	"fmt"
	"os"

	common "doppie.com/aoc-common"
)

func splitBeams(c []string) int {
	splitCount := 0

	// convert to byte for better modification
	grid := make([][]byte, len(c))
    for i := range c {
        grid[i] = []byte(c[i])
    }

	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i - 1][j] == 'S' {
				grid[i][j] = '|'
			}
			if grid[i - 1][j] == '|' {
				if grid[i][j] == '.' {
					grid[i][j] = '|'
				} else if grid[i][j] == '^' {
					didSplit := false
					if grid[i][j - 1] == '.' {
						grid[i][j - 1] = '|'
						didSplit = true
					}
					if grid[i][j + 1] == '.' {
						grid[i][j + 1] = '|'
						didSplit = true
					}
					if didSplit {
						splitCount++
					}
				}
			}
		}
	}

	return splitCount
}



func startQuantum(c []string) int {
	grid := make([][]byte, len(c))
    for i := range c {
        grid[i] = []byte(c[i])
    }

	sCol := 0

	for i := 0; i < len(c[0]); i++ {
		if c[0][i] == 'S' {
			sCol = i
		}
	}

	memo := make(map[[2]int]int)

	var quantumSplitBeams func(g [][]byte, row, col int) int
	quantumSplitBeams = func(g [][]byte, row, col int) int {
		val, ok := memo[[2]int{row, col}]
		if ok {
			return val
		}
		sum := 0

		l := len(g)

		notBroken := true

		for i := row; i < l; i++ {
			if (g[i][col] == '^') {
				sum += quantumSplitBeams(g, i, col - 1) + quantumSplitBeams(g, i, col + 1)
				notBroken = false
				break
			}
		}

		if notBroken {
			sum += 1
		}

		memo[[2]int{row, col}] = sum

		return sum
	}

	return quantumSplitBeams(grid, 1, sCol)
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)[1:]

	c := splitBeams(contents)

	fmt.Printf("The beam has been split %d times!\n", c)

	q := startQuantum(contents)

	fmt.Printf("The beam has been split into %d timelines!\n", q)
}