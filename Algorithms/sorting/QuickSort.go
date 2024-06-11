package sorting

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr)/2

	pivot := arr[mid]
	var right []int
	var left []int

	for i := 0; i < len(arr); i++ {
		if i == mid {
			continue
		}

		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	leftSorted := QuickSort(left)
	rightSorted := QuickSort(right)

	leftSorted = append(leftSorted, pivot)

	leftSorted = append(leftSorted, rightSorted...)

	return leftSorted
}