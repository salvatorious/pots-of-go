package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	memo "github.com/salvatorious/pots-of-go/memo"
	types "github.com/salvatorious/pots-of-go/types"
)

var optiCount int

func main() {
	readInput()
	cases := readInput()

	for _, tCase := range cases {
		if len(tCase) > 1 {
			simple := simpleSolve(tCase)
			optimal := optimalSolve(tCase)

			// if err != nil {
			// 	panic(err)
			// }

			fmt.Println("Simple: ", simple)
			fmt.Println("Optimal: ", optimal)
		}
		fmt.Printf("optimalSolve called %s times \n", strconv.Itoa(optiCount))
		fmt.Println("-------------------")
		optiCount = 0
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

func simpleSolve(set []int) types.Solution {
	left := set[0]
	right := set[len(set)-1]
	var pick = 0
	var newSet []int
	var remainderSolution types.Solution

	if left > right {
		newSet = set[1:]
		pick = set[0]
	} else {
		newSet = set[:len(set)-1]
		pick = set[len(set)-1]
	}

	if len(set) == 1 { // the last pick
		remainderSolution = types.Solution{
			FirstPlayerGold:  0,
			SecondPlayerGold: 0,
		}
	} else {
		remainderSolution = simpleSolve(newSet)
	}

	return types.Solution{
		FirstPlayerGold:  pick + remainderSolution.SecondPlayerGold,
		SecondPlayerGold: remainderSolution.FirstPlayerGold,
	}
}

/*

The optimal choice is:

max(choose_left, choose_right)
	where
		choose_left = gold_in_left_pot + optimalSolve(pots_without_left).SecondPlayerGold
		choose_right = gold_in_right_pot + optimalSolve(pots_without_right).SecondPlayerGold
*/

func optimalSolve(set []int) types.Solution {
	optiCount++
	m := memo.New(optimalSolve)

	var pick = 0
	var newSet []int
	var remainderSolution types.Solution

	n := len(set)

	leftPot := set[0]
	rightPot := set[n-1]

	if n == 1 {
		return types.Solution{
			FirstPlayerGold:  leftPot,
			SecondPlayerGold: 0,
		}
	} else if n == 2 {
		return types.Solution{
			FirstPlayerGold:  max(leftPot, rightPot),
			SecondPlayerGold: min(leftPot, rightPot),
		}
	} else {
		if (leftPot + m.Get(set[1:]).SecondPlayerGold) > (rightPot + m.Get(set[:n-1]).SecondPlayerGold) {
			pick = leftPot
			newSet = set[1:]
		} else {
			pick = rightPot
			newSet = set[:n-1]
		}
		// fmt.Println("value chosen is: ", pick)
		// fmt.Println("--")
	}

	remainderSolution = m.Get(newSet)

	return types.Solution{
		FirstPlayerGold:  pick + remainderSolution.SecondPlayerGold,
		SecondPlayerGold: remainderSolution.FirstPlayerGold,
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
