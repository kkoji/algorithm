package sort

func BubbleSort(numbers []int) []int {
	len_numbers := len(numbers)
	for i := 0; i < len_numbers; i++ {
		limit := len_numbers - 1 - i
		for j := 0; j < limit; j++ {
		 	if numbers[j] > numbers[j+1] {
		 		numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
		 	}
		}
	}

	return numbers
}
