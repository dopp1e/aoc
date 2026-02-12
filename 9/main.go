package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	common "doppie.com/aoc-common"
)

type areaSize struct {
	x1 int
	y1 int
	x2 int
	y2 int
	area int
}

type areaSizeSort []areaSize

func (a areaSizeSort) Len() int           { return len(a) }
func (a areaSizeSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a areaSizeSort) Less(i, j int) bool { return a[i].area < a[j].area }

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

const BORDER = -1
const UNASSIGNED = 0
const INSIDE = 1

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

func findBiggestRectangle(c []string) (int, []areaSize) {
	biggestArea := 0
	sizes := make([]areaSize, 0)

	for i := 0; i < len(c); i++ {
		for j := i + 1; j < len(c); j++ {
			x1, y1 := unpackCoordinates(c[i])
			x2, y2 := unpackCoordinates(c[j])

			x := math.Abs(float64(x2 - x1)) + 1
			y := math.Abs(float64(y2 - y1)) + 1
			area := int(x * y)
			sizes = append(sizes, areaSize{x1: x1, y1: y1, x2: x2, y2: y2, area: area})
			if area > biggestArea {
				//fmt.Printf("New biggest rectangle found between (%d, %d) and (%d, %d).\n", x1, y1, x2, y2)
				biggestArea = area
			}
		}
	}

	return biggestArea, sizes
}

func prepareFigure(figure [][]int8, x, y int) [][]int8 {
	// since all expansions are always done on the entire thing, we can assume the number of columns
	// of the first row is the same, as all rows following it
	for i := 0; i < len(figure); i++ {
		xDiff := (x + 5) - len(figure[i])
		if (xDiff > 0) {
			figure[i] = append(figure[i], make([]int8, xDiff)...)
		}
	}
	// add new rows with the current width
	yDiff := (y + 5) - len(figure)
	if yDiff > 0 {
		newRows := make([][]int8, yDiff)
		for i := 0; i < yDiff; i++ {
			newRows[i] = make([]int8, x + 5)
		}
		figure = append(figure, newRows...)
	}
	return figure
}

func fillMap(c []string, figure [][]int8) ([][]int8, int, int) {
	sizeY := 0
	sizeX := 0

	for i := 0; i < len(c); i++ {
		nextIndex := i + 1
		if nextIndex == len(c) {
			nextIndex = 0
		}
		x1, y1 := unpackCoordinates(c[i])
		if x1 > sizeX {
			sizeX = x1 + 2
		}
		if y1 > sizeY {
			sizeY = y1 + 2
		}
		x2, y2 := unpackCoordinates(c[nextIndex])
		start := min(x1, x2)
		end := max(x1, x2)
		if x1 == x2 {
			start = min(y1, y2)
			end = max(y1, y2)
		}
		figure = prepareFigure(figure, max(x1, x2), max(y1, y2))
		for a := start; a <= end; a++ {
			if x1 == x2 {
				figure[a][x1] = BORDER
			} else {
				//fmt.Printf("%d, %d - %d, %d - %d, %d\n", y1, a, sizeY, sizeX, len(figure), len(figure[0]))
				figure[y1][a] = BORDER
			}
		}
	}

	return figure, sizeX, sizeY
}

func printFigure(figure [][]int, sizeX, sizeY int) {
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			val := figure[y][x]
			if val == INSIDE {
				fmt.Print("X")
			} else if val == BORDER {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}
}

func isBorder(figure [][]int8, x, y int) bool {
	return figure[y][x] == BORDER
}

func isUnassigned(figure [][]int8, x, y int) bool {
	return figure[y][x] == UNASSIGNED
}

func isInside(figure [][]int8, x, y int) bool {
	return figure[y][x] == INSIDE
}

func set(figure [][]int8, x, y int) {
	figure[y][x] = INSIDE
}

func isOutsideFigure(figure [][]int8, x1, y1, x2, y2 int) bool {
	// normalize coordinates so iteration works regardless of order
	startX := min(x1, x2)
	endX := max(x1, x2)
	startY := min(y1, y2)
	endY := max(y1, y2)

	// for x := startX; x <= endX; x++ {
	// 	for y := startY; y <= endY; y++ {
	// 		if figure[y][x] == UNASSIGNED {
	// 			return true
	// 		}
	// 	}
	// }

	for x := startX; x <= endX; x++ {
		if figure[startY][x] == UNASSIGNED || figure[endY][x] == UNASSIGNED {
			return true
		}
	}

	for y := startY; y <= endY; y++ {
		if figure[y][startX] == UNASSIGNED || figure[y][endY] == UNASSIGNED {
			return true
		}
	}

	return false
}

func recursiveScanFill(figure [][]int8, x, y, sizeX, sizeY int) {
	if !isUnassigned(figure, x, y) {
		return
	}

	x1 := x;
	for x1 < sizeX && isUnassigned(figure, x1, y) {
		set(figure, x1, y)
		x1++
	}
	x1 = x - 1
	for x1 >= 0 && isUnassigned(figure, x1, y) {
		set(figure, x1, y)
		x1--
	}

	// test line above
	x1 = x
	for x1 < sizeX && isInside(figure, x1, y) {
		if y > 0 && isUnassigned(figure, x1, y - 1) {
			recursiveScanFill(figure, x1, y - 1, sizeX, sizeY)
		}
		x1++
	}
	x1 = x
	for x1 >= 0 && isInside(figure, x1, y) {
		if y > 0 && isUnassigned(figure, x1, y - 1) {
			recursiveScanFill(figure, x1, y - 1, sizeX, sizeY)
		}
		x1--
	}

	// test line below
	x1 = x
	for x1 < sizeX && isInside(figure, x1, y) {
		if y < sizeY - 1 && isUnassigned(figure, x1, y + 1) {
			recursiveScanFill(figure, x1, y + 1, sizeX, sizeY)
		}
		x1++
	}
	x1 = x - 1
	for x1 >= 0 && isInside(figure, x1, y) {
		if y < sizeY - 1 && isUnassigned(figure, x1, y + 1) {
			recursiveScanFill(figure, x1, y + 1, sizeX, sizeY)
		}
		x1--
	}
}

func findBiggestContainedRectangle(c []string, sizes []areaSize) int {
	start := time.Now()
	figure := make([][]int8, 1)
	figure, sizeX, sizeY := fillMap(c, figure)

	timeTrack(start, "fillMap")

	x1, y1 := unpackCoordinates(c[0])
	x2, y2 := unpackCoordinates(c[1])
	x3, y3 := unpackCoordinates(c[2])

	diffX := x2 - x1
	if diffX == 0 {
		diffX = x3 - x2
	}
	modX := diffX / int(math.Abs(float64(diffX)))
	diffY := y2 - y1
	if diffY == 0 {
		diffY = y3 - y2
	}
	modY := diffY / int(math.Abs(float64(diffY)))

	fmt.Printf("%d, %d\n", modX, modY)
	fmt.Printf("%d, %d\n", sizeX, sizeY)

	recursiveScanFill(figure, x1 + modX, y1 + modY, sizeX, sizeY)
	timeTrack(start, "scanFill")

	for i := 0; i < len(sizes) - 1; i++ {
		id := len(sizes) - 1 - i
		
		if isOutsideFigure(figure, sizes[id].x1, sizes[id].y1, sizes[id].x2, sizes[id].y2) {
			continue
		}
		timeTrack(start, "findArea")
		return sizes[id].area
	}

	return 0
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)[1:]

	c, sizes := findBiggestRectangle(contents)

	sort.Sort(areaSizeSort(sizes))

	fmt.Printf("The size of the biggest rectangle equals %d.\n", c)

	d := findBiggestContainedRectangle(contents, sizes)

	fmt.Printf("The size of the biggest contained rectangle equals %d.\n", d)
}