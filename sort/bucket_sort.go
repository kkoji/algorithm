package sort

func BucketSort(numbers []int) []int {
	// 数列の最大値
	max_num := max(numbers)
	// 配列の要素数
	len_numbers := len(numbers)
	// バケットサイズ
	bucket_size := max_num / len_numbers
	// バケット
	buckets := make([][]int, bucket_size)

	for _, number := range numbers {
		// バケットサイズで割った数をインデックスとする
		i := number / bucket_size
		if i < bucket_size {
			// インデックスがバケットサイズ未満の場合は対象インデックスのバケットに格納
			buckets[i] = append(buckets[i], number)
		} else {
			// バケットサイズを超えるインデックスの場合は最後のバケットに格納
			buckets[bucket_size-1] = append(buckets[bucket_size-1], number)
		}
	}

	idx := 0
	for _, bucket := range buckets {
		// 各バケットごとにソートした結果をnumbersに格納していく
		for _, n := range InsertionSort(bucket) {
			numbers[idx] = n
			idx++
		}
	}

	return numbers
}

func max(numbers []int) int {
	tmp_max := 0
	for i := 0; i < len(numbers); i++ {
		if tmp_max < numbers[i] {
			tmp_max = numbers[i]
		}
	}
	return tmp_max
}
