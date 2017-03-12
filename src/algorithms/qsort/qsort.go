package qsort

func quickSort(values []int, left, right int) {

	k := left
	i, j := left, right
	for i < j {
		for j > k && values[j] >= values[k] {
			j--
		}
		values[j], values[k] = values[k], values[j]
		k = j
		for i < k && values[i] <= values[k] {
			i++
		}
		values[i], values[k] = values[k], values[i]
		k = i
	}
	if k-left > 1 {
		quickSort(values, left, k-1)
	}
	if right-k > 1 {
		quickSort(values, k+1, right)
	}
}

func QuickSort(values []int) {
	left := 0
	right := len(values) - 1
	quickSort(values, left, right)
}
