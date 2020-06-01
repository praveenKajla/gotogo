package arrays

import "my.app/utils"

//FindDuplicateHashing :
func FindDuplicateHashing(arr []int) int {

	m := make(map[int]bool)

	for i := 0; i < len(arr); i++ {
		if _, ok := m[arr[i]]; ok {
			return arr[i]
		}
		m[arr[i]] = true
	}
	return -1

}

//FindDuplicateAbs : https://www.youtube.com/watch?v=HuZJqRDOPo0
func FindDuplicateAbs(arr []int) int {

	duplicate := -1

	for i := 0; i < len(arr); i++ {
		val := arr[utils.Abs(arr[i])]
		if val >= 0 {
			arr[utils.Abs(arr[i])] = -arr[utils.Abs(arr[i])]
		} else {
			duplicate = utils.Abs(arr[i])
			break
		}
	}
	//convert back to original array
	for  i := 0; i < len(arr); i++ {
		arr[i] = utils.Abs(arr[i])
	}
	return duplicate

}
