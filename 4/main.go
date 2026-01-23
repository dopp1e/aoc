package main

import (
	"fmt"
	"os"
	"unicode/utf8"

	"doppie.com/aoc-common"
)

func getGridSize(c []string) (int, int) {
    if len(c) == 0 {
        return 0, 0
    }

    y := len(c)
    x := utf8.RuneCountInString(c[1])
    fmt.Printf("(%d, %d)", x, y)
    return x, y
}

func areCoordsWithinBounds(x, y, sizeX, sizeY int) bool {
    return (x >= 0 && y >= 0 && x < sizeX && y < sizeY)
}

func canBeAccessed(contents []string, x, y, sizeX, sizeY int) bool {
    neighbors := 0

    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            cx := x + i
            cy := y + j
            if cx == x && cy == y {
                continue
            }
            if areCoordsWithinBounds(cx, cy, sizeX, sizeY) {
                if contents[cx][cy] == '@' {
                    neighbors++
                }
            }
        }
    }

    fmt.Printf("At (%d, %d) found %d neighbors.\n", x, y, neighbors)

    return neighbors < 4
}

func findPaperRolls(contents []string) int {
    sum := 0

    x, y := getGridSize(contents)

    if x == 0 || y == 0 {
        return 0
    }

    for i := range x {
        for j := range y {
            if contents[i][j] == '@' && canBeAccessed(contents, i, j, x, y) {
                sum++
            }
        }
    }

    return sum
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)[1:]

    rolls := findPaperRolls(contents)

    fmt.Printf("A forklift can access %d paper rolls.\n", rolls)
}
