package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	common "doppie.com/aoc-common"
)

func unpackCoordinates(s string) (int, int) {
	spl := strings.Split(s, ",")
	x, err := strconv.Atoi(spl[0])
	if err != nil {
		log.Fatalf("Failed converting first element of %s to int!\n", s)
	}
	y, err := strconv.Atoi(spl[1])
	if err != nil {
		log.Fatalf("Failed converting second element of %s to int!\n", s)
	}
	return x, y
}

func findBiggestRectangle(c []string) int {
	biggestArea := 0

	for i := 0; i < len(c); i++ {
		for j := i + 1; j < len(c); j++ {
			x1, y1 := unpackCoordinates(c[i])
			x2, y2 := unpackCoordinates(c[j])

			x := math.Abs(float64(x2 - x1)) + 1
			y := math.Abs(float64(y2 - y1)) + 1
			area := int(x * y)
			if area > biggestArea {
				//fmt.Printf("New biggest rectangle found between (%d, %d) and (%d, %d).\n", x1, y1, x2, y2)
				biggestArea = area
			}
		}
	}

	return biggestArea
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)[1:]

	c := findBiggestRectangle(contents)

	fmt.Printf("The size of the biggest rectangle equals %d.\n", c)
}