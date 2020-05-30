package dp

import (
	"my.app/utils"
)

//LIS : https://youtu.be/Ns4LCeeOFS4
func LIS(inp []int) int {
	// input : {3, 10, 2, 1, 20}
	// LIS: {3,10,20} => Longest increasing subsequence
	// taking each index as starting point we calculate the lis

	// startting from last element i.e. 3 {10,2,1,20}

	lis := make([]int, len(inp))
	lis[0] = 1
	for i := 1; i < len(inp); i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if inp[j] < inp[i] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
	}
	return utils.MaxIntArray(lis)

}
