package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	common "doppie.com/aoc-common"
)

func getLights(s string) []bool {
	l := len(s) - 2
	lights := make([]bool, l)
	for i := 0; i < l; i++ {
		id := i + 1
		if s[id] == '#' {
			lights[i] = true
		}
	}

	return lights
}

func getWirings(s []string) [][]int {
	l := len(s)
	wirings := make([][]int, l)
	for i := 0; i < l; i++ {
		spl := strings.Split(s[i][1:len(s[i]) - 1], ",")
		v := make([]int, 0)
		if len(spl) == 1 {
			str := string(s[i][1])
			n, err := strconv.Atoi(str)
			if err != nil {
				log.Fatalf("Conversion (wire) of %s failed!", str)
			}
			v = append(v, n)
		} else {
			for _, w := range spl {
				n, err := strconv.Atoi(w)
				if err != nil {
					log.Fatalf("Conversion (wire) of %s failed!", w)
				}
				v = append(v, n)
			}
		}
		wirings[i] = v
	}

	return wirings
}

func getJoltage(s string) []int {
	spl := strings.Split(s[1:len(s)-1], ",")
	joltages := make([]int, len(spl))
	for i, j := range spl {
		n, err := strconv.Atoi(j)
		if err != nil {
			log.Fatalf("Conversion (jolt) of %s failed!", j)
		}
		joltages[i] = n
	}

	return joltages
}

func getMachineDetails(s string) ([]bool, [][]int, []int) {
	spl := strings.Split(s, " ")
	lightsString := spl[0]
	wiringStrings := spl[1:len(spl)-1]
	joltageString := spl[len(spl)-1]

	lights := getLights(lightsString)
	wirings := getWirings(wiringStrings)
	joltage := getJoltage(joltageString)

	// for _, l := range lights {
	// 	fmt.Printf("%t ", l)
	// }
	// fmt.Print("\n")
	// for _, w := range wirings {
	// 	for _, ww := range w {
	// 		fmt.Printf("%d ", ww)
	// 	}
	// 	fmt.Print("\n")
	// }
	// for _, j := range joltage {
	// 	fmt.Printf("%d ", j)
	// }
	// fmt.Print("\n\n")

	return lights, wirings, joltage
}

func factorial(x int) int {
	n := 1
	if x == 0 {
		return n
	}
	return x * factorial(x - 1)
}

// given a list containing lists of numbers,
// return a list of all possible joins between 1-n of the lists
func permute(w [][]int) [][]int {
	var res [][]int
	for i := 1; i <= len(w); i++ {
		res = append(res, permuteHelper(w, i)...)
	}
	return res
}

func permuteHelper(w [][]int, n int) [][]int {
	if n > len(w) {
		return nil
	}
	if n == 1 {
		res := make([][]int, len(w))
		for i := range w {
			cp := make([]int, len(w[i]))
			copy(cp, w[i])
			res[i] = cp
		}
		return res
	}

	var res [][]int
	for i := 0; i < len(w); i++ {
		for _, p := range permuteHelper(w[i+1:], n-1) {
			cp := make([]int, len(w[i])+len(p))
			copy(cp, w[i])
			copy(cp[len(w[i]):], p)
			res = append(res, cp)
		}
	}

	return res
}

func findCombo(l []bool, w [][]int) int {
	for count := 1; count <= len(w); count++ {
		perms := permuteHelper(w, count)
		//fmt.Printf("%d - %d - %d\n", count, len(perms), len(w))
		for _, p := range perms {
			if wouldBeEnabled(make([]bool, len(l)), l, p) {
				return count
			}
		}
	}

	return -1
}

func wouldBeEnabled(l, target []bool, w []int) bool {
	// operate on a copy so caller's slice isn't mutated
	cur := make([]bool, len(l))
	copy(cur, l)
	//fmt.Print("New perm:\n")
	for i := 0; i < len(w); i++ {
		idx := w[i]
		//fmt.Printf("%d\n", idx)
		cur[idx] = !cur[idx]
	}
	// for i := 0; i < len(cur); i++ {
	// 	fmt.Printf("%t ", cur[i])
	// }
	// fmt.Print("\n")
	for i := 0; i < len(cur); i++ {
		if cur[i] != target[i] {
			return false
		}
	}
	//fmt.Print("Enabled!\n")
	return true
}

func enableMachines(c []string) int {
	sum := 0

	for _, m := range c {
		l, w, _ := getMachineDetails(m)
		v := findCombo(l, w)
		fmt.Printf("Value: %d\n", v)
		sum += v
	}

	return sum
}


func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)[1:]

	sum := enableMachines(contents)
	log.Printf("Sum: %d", sum)
}