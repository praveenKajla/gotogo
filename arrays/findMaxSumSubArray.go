package arrays


//Given an array of integers, find largest sub-array formed by consecutive integers. The sub-array should contain all distinct values.

//FindMaxSumSubArray : kadane's algo : https://www.youtube.com/watch?v=86CQq3pKSUw
func FindMaxSumSubArray(arr []int) int{

	n := len(arr)
	if n==1 {
		return arr[0]
	}

	var  maxSum int = arr[0]
	maxSoFar := arr[0]

	for i:= 1;i<n;i++{
		
		if maxSoFar + arr[i] < arr[i] {
			maxSoFar = arr[i]
		} else {
			maxSoFar = maxSoFar + arr[i]
		}
		if maxSoFar > maxSum {
			maxSum = maxSoFar
		}
	}
	return maxSum

}

