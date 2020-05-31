package arrays

import (
	"fmt"
	"sort"
)

// Given an unsorted array of integers, find a pair with given sum in it : O(n2)
// https://www.techiedelight.com/find-pair-with-given-sum-array/

//FindPairNaive : On(2)
func FindPairNaive(arr []int, sum int) {

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				fmt.Println("Pair found at index ", i, "and ", j)
				return
			}

		}

	}
}

//FindPairSorting :  O(nlogn)
func FindPairSorting(arr []int, sum int) {
	sort.Ints(arr)

	var high int = len(arr) - 1
	var low int = 0

	for low < high {

		if arr[high]+arr[low] == sum {
			fmt.Println("Pair found at index ", low, "and ", high)
			return
		}
		if arr[high]+arr[low] < sum {
			low++
		} else {
			high--
		}

	}
	fmt.Println("Pair  not found")

}

//FindPairHashmap : O(n)
func FindPairHashmap(arr []int, sum int) {

	m := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if val, ok := m[sum-arr[i]]; ok {
			fmt.Println("Pair found at index", val, "and", i)
			return
		}
		m[arr[i]] = i
	}
	fmt.Println("Pair not found")
}
