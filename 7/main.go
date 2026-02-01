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

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)[1:]

	c := splitBeams(contents)

	fmt.Printf("The beam has been split %d times!", c)

}