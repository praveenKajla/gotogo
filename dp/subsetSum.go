package dp

func SubsetSum(arr []int, sum int) bool {

	return SubsetSumHelper(arr, sum, len(arr))

}

func SubsetSumHelper(arr []int, sum int, n int) bool {

	if sum == 0 {
		return true
	}
	if n == 0 && sum != 0 {
		return false
	}
	if arr[n-1] > sum {
		return SubsetSumHelper(arr, sum, n-1)
	}
	return SubsetSumHelper(arr, sum, n-1) || SubsetSumHelper(arr, sum-arr[n-1], n-1)
}
