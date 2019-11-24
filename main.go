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
			fmt.Println("Optimal: ", optimalSolve(tCase))
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

/*

The optimal choice is:

max(choose_left, choose_right)
	where
		choose_left = gold_in_left_pot + optimalSolve(pots_without_left).secondPlayerGold
		choose_right = gold_in_right_pot + optimalSolve(pots_without_right).secondPlayerGold
*/
func optimalSolve(set []int) solution {

	var pick = 0
	var newSet []int
	var remainderSolution solution

	n := len(set)
	leftPot := set[0]
	rightPot := set[n-1]

	if n == 1 {
		return solution{
			firstPlayerGold:  leftPot,
			secondPlayerGold: 0,
		}
	} else if n == 2 {
		return solution{
			firstPlayerGold:  max(leftPot, rightPot),
			secondPlayerGold: min(leftPot, rightPot),
		}
	} else {
		if (leftPot + optimalSolve(set[1:]).secondPlayerGold) > (rightPot + optimalSolve(set[:n-1]).secondPlayerGold) {
			pick = leftPot
			newSet = set[1:]
		} else {
			pick = rightPot
			newSet = set[:n-1]
		}
		// fmt.Println("value chosen is: ", pick)
		// fmt.Println("--")
	}

	remainderSolution = optimalSolve(newSet)

	return solution{
		firstPlayerGold:  pick + remainderSolution.secondPlayerGold,
		secondPlayerGold: remainderSolution.firstPlayerGold,
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
