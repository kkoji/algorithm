package sort

func ShellSort(numbers []int) []int {
	len_numbers := len(numbers)
	gap := len_numbers / 2

	// gapが0になったらソート完了
	for gap > 0 {
		for i := 0; i < len_numbers; i++ {
			// 後ろの比較対象の値
			temp := numbers[i]
			j := i

			// 手前の比較対象のindexが0以上、かつ手前の値の方が大きい場合
			for j-gap >= 0 && numbers[j-gap] > temp {
				// 手前の値を後ろの位置に代入
				numbers[j] = numbers[j-gap]
				// ギャップ分手前にずらす
				j -= gap
			}
			// 後ろの値を手前の位置に代入
			numbers[j] = temp
		}
		gap /= 2
	}

	return numbers
}
