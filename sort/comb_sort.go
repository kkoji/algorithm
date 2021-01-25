package sort

func CombSort(numbers []int) []int {
	len_numbers := len(numbers)
	gap := len_numbers
	swapped := false

	for gap != 1 || swapped {
		gap = int(float32(gap) / 1.3)

		if gap < 1 {
			gap = 1
		}

		swapped = false
		for i := 0; i < (len_numbers - gap); i++ {
			if numbers[i] > numbers[i+gap] {
				numbers[i], numbers[i+gap] = numbers[i+gap], numbers[i]
				swapped = true
			}
		}
	}

	return numbers
}
