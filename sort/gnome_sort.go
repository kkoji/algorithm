package sort

func GnomeSort(numbers []int) []int {
	len_numbers := len(numbers)
	index := 0

	for index < len_numbers {
		if index == 0 {
			index += 1
		}
		if numbers[index] >= numbers[index-1] {
			index += 1
		} else {
			numbers[index], numbers[index-1] = numbers[index-1], numbers[index]
			index -= 1
		}
	}

	return numbers
}
