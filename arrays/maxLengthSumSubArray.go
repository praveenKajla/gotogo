package arrays

// Given an array of integers, find maximum length sub-array having given sum.
// https://www.youtube.com/watch?v=HJDlxZNe1UI

//MaxLengthSumSubArray : O(n) solution
func MaxLengthSumSubArray(arr []int , sum int)[]int {

	m := make(map[int]int)

	length := 0
	endIndex := -1


	sumSoFar := 0

	for i:=0;i<len(arr);i++{
		sumSoFar += arr[i]

		if sumSoFar == sum {
			return arr[0:i+1]
		}

		// bcoz we need the max length so if sum repeats we need the whole length thats why !ok
		if _,ok := m[sumSoFar];!ok {
			m[sumSoFar] = i
		}

		if val,ok := m[sumSoFar-sum];ok && length < i-val {
			length = i - val
			endIndex = i

		}
	}
	if endIndex > 0 {
		return  arr[endIndex - length + 1 : endIndex+1]
	} 
	return make([]int, 0)
	

}