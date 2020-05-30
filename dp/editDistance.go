package dp

import (
	"math"
)

// Given two strings str1 and str2 and below operations that can performed on str1.
// Find minimum number of edits (operations) required to convert ‘str1’ into ‘str2’.
// Insert
// Remove
// Replace

//eg : M sunday -> N saturday
// y == y => minEdits(M(0,4),N(0,6)) => a == a => minEdits(M(0,3),N(0,5)) => d == d => minEdits(M(0,2),N(0,4))
//i.e.  sun ->  satur => n != r => 1 + min(minEdits(M(0,1),N(0,3)),minEdits(M(0,1),N(0,4)),minEdits(M(0,2),N(0,3)
// i.e => 1 + min(replace,remove,insert)

//EditDistance :algo
func EditDistance(str1 []rune, str2 []rune) int {

	return EditDistanceHelper(str1, str2, len(str1), len(str2))

}

func EditDistanceHelper(str1 []rune, str2 []rune, m int, n int) int {

	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	if str1[m-1] == str2[n-1] {
		return EditDistanceHelper(str1, str2, m-1, n-1)
	}
	return 1 + int(math.Min(math.Min(float64(EditDistanceHelper(str1, str2, m, n-1)), float64(EditDistanceHelper(str1, str2, m-1, n))), float64(EditDistanceHelper(str1, str2, m-1, n-1))))

}

//EditDistanceMemo : dynamic programming , memoization
func EditDistanceMemo(str1 []rune, str2 []rune) int {
	m := len(str1)
	n := len(str2)

	memo := make([][]int, m+1)
	for row := range memo {
		memo[row] = make([]int, n+1)
	}

	for i := 0; i <= len(str1); i++ {
		for j := 0; j <= len(str2); j++ {
			if i == 0 {
				memo[i][j] = j
			} else if j == 0 {
				memo[i][j] = i
			} else if str1[i-1] == str2[j-1] {
				memo[i][j] = memo[i-1][j-1]
			} else {
				memo[i][j] = 1 + int(math.Min(math.Min(float64(memo[i-1][j]), float64(memo[i][j-1])), float64(memo[i-1][j-1])))
			}

		}
	}
	return memo[m][n]
}
