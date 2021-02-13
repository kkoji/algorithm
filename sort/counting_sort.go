package sort

func CountingSort(numbers []int) []int {
	max_num := max(numbers)
	// 各数字の出現回数をカウントするための変数
	counts := make([]int, max_num+1)
	// ソート結果
	result := make([]int, len(numbers))

	//カウント処理
	for _, num := range numbers {
		counts[num] += 1
	}
	// 先頭からのカウント数累計をそれぞれの位置に格納する
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	// 数列末尾のインデックス
	i := len(numbers) - 1
	for i >= 0 {
		index := numbers[i]
		// 対象位置に数字を格納
		result[counts[index]-1] = numbers[i]
		// 格納先indexをスライド（同じ番号が出現した場合に左隣に格納するため）
		counts[index] -= 1
		i -= 1
	}

	return result
}
