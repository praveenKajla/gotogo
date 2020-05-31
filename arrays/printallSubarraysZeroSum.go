package arrays

import "fmt"

// Given an array of integers, print all sub-arrays with 0 sum.
// https://www.youtube.com/watch?v=hLcYp67wCcM

//PrintallSubarraysZeroSumNaive : O(n2)
func PrintallSubarraysZeroSumNaive(arr []int) {
	n := len(arr)

	var sum int = 0

	for i := 0; i < n; i++ {
		sum = 0
		for j := i; j < n; j++ {
			sum += arr[j]

			if sum == 0 {
				fmt.Println("Subarry with sum 0 :  ", arr[i:j])
			}

		}
	}
}

//PrintallSubarraysZeroSum : o(n)
func PrintallSubarraysZeroSum(arr []int) {

	m := make(map[int][]int)
	var sum int = 0

	for i := 0; i <= len(arr); i++ {
		if val, ok := m[sum]; ok {
			for _, index := range val {
				fmt.Println("Subarry with sum 0 :  ", index, i-1, arr[index:i])
			}

		}
		if i == len(arr) {
			return
		}
		m[sum] = append(m[sum], i) //important
		sum += arr[i]
	}
}
