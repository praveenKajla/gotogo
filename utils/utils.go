package utils

func MaxIntArray(array []int) int {
	var max int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}


func Abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}