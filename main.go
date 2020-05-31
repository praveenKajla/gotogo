package main

import (
	"fmt"

	"my.app/arrays"
	"my.app/dp"
	"my.app/tree"
)

func main() {
	// fmt.Println(dp.FibonacciRecursive(10))
	// fmt.Println(dp.FibonacciIterative(10))
	// fmt.Println(dp.FibonacciFast(10))
	// coins := []int{1, 2, 5}
	// fmt.Println(dp.BruteForceCoinsChange(coins[0:], 11))
	// fmt.Println(dp.MinCoinsChange(coins[0:], 11))
	// fmt.Println(dp.MinCoinsChangeMemo(coins[0:], 11))
	// fmt.Println(dp.MinCoinsChangeMemoIterative(coins[0:], 11))
	// fmt.Println(dp.PrintLCS([]rune("AGGTAB"), []rune("GXTXAYB")))
	// fmt.Println(dp.LIS([]int{3, 10, 2, 1, 20}))
	// fmt.Println(dp.EditDistance([]rune("sunday"), []rune("saturday")))
	// fmt.Println(dp.EditDistanceMemo([]rune("sunday"), []rune("saturday")))
	fmt.Println(dp.SubsetSum([]int{3, 34, 4, 12, 5, 2}, 9))
	tree.BinaryTreeExample()
	arrays.FindPairHashmap([]int{8, 7, 2, 5, 3, 1}, 10)
	// arrays.PrintallSubarraysZeroSumNaive([] int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2})
	arrays.PrintallSubarraysZeroSum([] int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2})

}
