package sort

func SelectionSort(numbers []int) []int {
	len_numbers := len(numbers)

	for i := 0; i < len_numbers; i++ {
		min_idx := i
		for j := i+1; j < len_numbers; j++ {
			if numbers[min_idx] > numbers[j] {
				min_idx = j
			}
		}
		numbers[i], numbers[min_idx] = numbers[min_idx], numbers[i]
	}

	return numbers
}
