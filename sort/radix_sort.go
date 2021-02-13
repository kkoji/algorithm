package sort

func countingSort(numbers []int, place int) []int {
	// 各数字の出現回数をカウントするための変数
	counts := make([]int, 10)
	// ソート結果
	result := make([]int, len(numbers))

	//カウント処理
	for _, num := range numbers {
		index := num / place % 10
		counts[index] += 1
	}
	// 先頭からのカウント数累計をそれぞれの位置に格納する
	for i := 1; i < 10; i++ {
		counts[i] += counts[i-1]
	}

	// 数列末尾のインデックス
	i := len(numbers) - 1
	for i >= 0 {
		index := numbers[i] / place % 10
		// 対象位置に数字を格納
		result[counts[index]-1] = numbers[i]
		// 格納先indexをスライド（同じ番号が出現した場合に左隣に格納するため）
		counts[index] -= 1
		i -= 1
	}

	return result
}

func RadixSort(numbers []int) []int {
	max_num := max(numbers)
	// ソート対象の位
	place := 1
	// 位ごとに（一桁目から順に）カウンティングソートを繰り返すのがラディックスソート
	for max_num > place {
		numbers = countingSort(numbers, place)
		place *= 10
	}

	return numbers
}
