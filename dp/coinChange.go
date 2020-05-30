package dp

import (
	"math"
)

//https://www.youtube.com/watch?v=1R0_7HqNaW0

//BruteForceCoinsChange brute force solution
func BruteForceCoinsChange(coins []int, amount int) int {
	// [1,2,3]

	return helper(coins, len(coins), amount)

}

func helper(coins []int, m int, amount int) int {

	if amount == 0 {
		return 1
	}

	if amount < 0 {
		return 0
	}
	//if no amount left or no coins available to include in current amount return 0
	if m <= 0 && amount >= 1 {
		return 0
	}
	// including mth coins and excluding mth coin
	return helper(coins, m, amount-coins[m-1]) + helper(coins, m-1, amount)

}

func MinCoinsChange(coins []int, amount int) int {
	return MinCoinsChangeHelper(coins, len(coins), amount)

}
func MinCoinsChangeHelper(coins []int, m int, amount int) int {
	//if no amount left i.e. base
	if amount == 0 {
		return 0
	}

	// if negative amount discard the result
	if amount < 0 {
		return -1
	}

	//let result be max integer for comparison
	res := math.MaxUint32

	for i := 0; i < m; i++ {
		subproblem := int(MinCoinsChangeHelper(coins, m, amount-coins[i]))
		if subproblem == -1 {
			continue
		}

		res = int(math.Min(float64(res), float64(1+subproblem)))

	}
	if res != math.MaxUint32 {
		return res
	}
	return -1
}

func MinCoinsChangeMemo(coins []int, amount int) int {
	memo := make(map[int]int)
	return MinCoinsChangeMemoHelper(coins, len(coins), amount, memo)

}
func MinCoinsChangeMemoHelper(coins []int, m int, amount int, memo map[int]int) int {

	//if already calculated
	if val, ok := memo[amount]; ok {
		return val
	}
	if amount == 0 {
		return 0
	}

	// if negative amount discard the result
	if amount < 0 {
		return -1
	}

	//let result be max integer for comparison
	res := math.MaxUint32

	for i := 0; i < m; i++ {
		subproblem := int(MinCoinsChangeMemoHelper(coins, m, amount-coins[i], memo))
		if subproblem == -1 {
			continue
		}

		res = int(math.Min(float64(res), float64(1+subproblem)))

	}
	if res != math.MaxUint32 {
		memo[amount] = res
		return memo[amount]
	}
	return -1
}

func MinCoinsChangeMemoIterative(coins []int, amount int) int {

	return MinCoinsChangeMemoHelperIterative(coins, len(coins), amount)

}
func MinCoinsChangeMemoHelperIterative(coins []int, m int, amount int) int {

	//if no legit amount

	// if negative amount discard the result
	if amount <= 0 {
		return -1
	}
	memo := make([]int, amount+1)
	for i := 1; i < len(memo); i++ {
		memo[i] = amount + 1

	}
	for i := 0; i < len(memo); i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}
			memo[i] = int(math.Min(float64(memo[i]), float64(1+memo[i-coins[j]])))

		}

	}
	if memo[amount] == amount+1 {
		return -1
	}
	return memo[amount]
}
