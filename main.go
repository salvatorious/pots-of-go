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
		if len(tCase) > 1 {
			fmt.Println("Simple: ", simpleSolve(tCase))
			fmt.Println("Optimal: ", optimalSolve(tCase, len(tCase)-1))
			fmt.Println("-------------------")
		}
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

// var memo [][]int

func optimalSolve(set []int, n int) solution {
	// fmt.Println("value of n is: ", n)
	// fmt.Println("--")

	var pick = 0
	var newSet []int
	var remainderSolution solution

	if set[0] == set[n] {
		pick = set[0]
		newSet = set[1:]
	} else if set[0] == n+1 {
		pick = max(set[0], set[n])

		if set[0] > set[n] {
			newSet = set[1:]
		} else {
			newSet = set[:n]
		}
	} else {
		if (set[0] + min(set[n], set[1])) > (set[n] + min(set[0], set[n-1])) {
			pick = set[0]
			newSet = set[1:]
		} else {
			pick = set[n]
			newSet = set[:n]
		}
	}

	if len(set) == 1 { // the last pick
		remainderSolution = solution{
			firstPlayerGold:  0,
			secondPlayerGold: 0,
		}
	} else {
		remainderSolution = optimalSolve(newSet, len(newSet)-1)
	}

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
