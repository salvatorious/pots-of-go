package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readInput()
	cases := readInput()

	for _, tCase := range cases {
		fmt.Println("Simple: ", simpleSolve(tCase))
		fmt.Println("Smarter: ", optimalSolve(tCase, len(tCase)-1))
		fmt.Println("-------------------")
	}
}

func readInput() [][]int {
	fileHandle, _ := os.Open("inputfile.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	var testCases [][]int

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), " ")

		var potSet []int
		for _, strNum := range line {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				panic(err)
			}
			potSet = append(potSet, num)
		}

		testCases = append(testCases, potSet)
	}

	return testCases
}

// func greedyPick(set []int) []int {
// 	left := set[0]
// 	lAdj := set[1]
// 	right := set[len(set)-1]
// 	rAdj := set[len(set)-2]
// 	var newSet []int

// 	if (left + min(right, lAdj)) > (right + min(left, rAdj)) {
// 		newSet = set[1:]
// 	} else {
// 		newSet = set[:len(set)-1]
// 	}

// 	return newSet
// }

type solution struct {
	firstPlayerGold  int
	secondPlayerGold int
}

func simpleSolve(set []int) solution {
	left := set[0]
	right := set[len(set)-1]
	var pick = 0
	var newSet []int
	var remainderSolution solution

	if left > right {
		newSet = set[1:]
		pick = set[0]
	} else {
		newSet = set[:len(set)-1]
		pick = set[len(set)-1]
	}

	if len(set) == 1 { // the last pick
		remainderSolution = solution{
			firstPlayerGold:  0,
			secondPlayerGold: 0,
		}
	} else {
		remainderSolution = simpleSolve(newSet)
	}

	return solution{
		firstPlayerGold:  pick + remainderSolution.secondPlayerGold,
		secondPlayerGold: remainderSolution.firstPlayerGold,
	}
}

// memo := [][]int

func optimalSolve(set []int, i int, j int) solution {
	left := set[0]
	right := set[len(set)-1]

	var pick = 0
	var newSet []int
	var remainderSolution solution

	// lAdj := set[1]
	// rAdj := set[len(set)-2]

	if (left + min(right, lAdj)) > (right + min(left, rAdj)) {
		newSet = set[1:]
		pick = set[0]
	} else {
		newSet = set[:len(set)-1]
		pick = set[len(set)-1]
	}

	remainderSolution = optimalSolve(newSet)

	return solution{
		firstPlayerGold:  pick + remainderSolution.secondPlayerGold,
		secondPlayerGold: remainderSolution.firstPlayerGold,
	}
}

// max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
