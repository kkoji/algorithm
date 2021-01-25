package sort

import (
	"math/rand"
)

// 昇順に並んでいるか検証する関数
func in_order(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}
	return true
}

// 適当に並び替えて、数字順になるまで繰り返すのがボゴソート
func BogoSort(numbers []int) []int {
	for !in_order(numbers) {
		rand.Shuffle(len(numbers), func(i int, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})
	}

	return numbers
}
