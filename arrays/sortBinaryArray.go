package arrays

import "fmt"

//Given a binary array, sort it in linear time and constant space. Output should print contain all zeroes followed by all ones.

//SortBinaryArrayNaive : O(n)
func SortBinaryArrayNaive(arr []int) []int {

	var zeroes int = 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zeroes++
		}
	}

	var k int = 0

	for zeroes >= 1 {
		zeroes--
		arr[k] = 0
		k++
	}
	for k < len(arr)-1 {
		k++
		arr[k] = 1
	}
	return arr

}

//SortBinaryArray :
func SortBinaryArray(arr []int) []int {

	var k int = 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[k] = 0
			k++
		}
	}

	for i := k; i < len(arr); i++ {
		arr[i] = 1
	}
	return arr

}

// SortBinaryArrayPartitioning :   { 0, 0, 1, 1, 1, 0, 0, 1 }
func SortBinaryArrayPartitioning(arr []int) []int {
	var pivot int = 1

	var pIndex int = 0

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr, pIndex)
		if arr[i] < pivot {
			arr[i], arr[pIndex] = arr[pIndex], arr[i]
			pIndex++

		}
	}
	return arr
}
