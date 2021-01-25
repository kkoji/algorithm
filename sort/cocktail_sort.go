package sort

func CocktailSort(numbers []int) []int {
	len_numbers := len(numbers)
	swapped := true
	start := 0
	end := len_numbers - 1

	for swapped {
		swapped = false
		for i := start; i < end; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		swapped = false
		end--
		for i := end; i > start; i-- {
			if numbers[i] < numbers[i-1] {
				numbers[i], numbers[i-1] = numbers[i-1], numbers[i]
				swapped = true
			}
		}
		start++
	}

	return numbers
}
