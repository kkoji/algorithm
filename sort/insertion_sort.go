package sort

func InsertionSort(numbers []int) []int {
	len_numbers := len(numbers)

	for i := 1; i < len_numbers; i++ {
		// 移動させる数字
		temp := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > temp {
			// 移動させる数字の方が小さい場合、大きい数字を右にずらす
			numbers[j+1] = numbers[j]
			j -= 1
		}
		// 最後に、移動させる数字を適切な位置に移動
		numbers[j+1] = temp
	}

	return numbers
}
