package arrays

//QuickSort : https://www.youtube.com/watch?v=ZHVk2blR45Q
func QuickSort(arr []int) []int {
	QuickSortHelper(arr, 0, len(arr)-1)
	return arr
}

func QuickSortHelper(arr []int, start int, end int) {

	if start >= end {
		return
	}

	pivot := Partition(arr, start, end)
	QuickSortHelper(arr, start, pivot-1)
	QuickSortHelper(arr, pivot+1, end)

}

func Partition(arr []int, start int, end int) int {
	var pivot int = arr[end]

	var pIndex int = start

	for i := start; i < end; i++ {

		if arr[i] <= pivot {
			arr[i], arr[pIndex] = arr[pIndex], arr[i]
			pIndex++

		}
	}
	arr[end], arr[pIndex] = arr[pIndex], arr[end]
	return pIndex
}
