package main

import (
	"math/rand"
	"time"
)

// ソート対象のランダムな数字を含むスライスを作成する
func generate_number_list(int_num int) []int {
	// ソート対象の配列を作成
	rand.Seed(time.Now().UnixNano())
	var numbers []int
	// 指定回数ランダムな数字を配列に追加
	for i := 0; i < int_num; i++ {
		// 1 - 1000までの範囲でランダムな数字を追加
		numbers = append(numbers, rand.Intn(100) + 1)
	}

	return numbers
}

func main() {
	// original_number_list := generate_number_list(10)
	// number_list := make([]int, len(original_number_list))
	// Bogo Sort
	// copy(number_list, original_number_list)
	// fmt.Println("Before: Bogo Sort", number_list)
	// fmt.Println("After Bogo Sort", sort.BogoSort(number_list))

	// Bubble Sort
	// copy(number_list, original_number_list)
	// fmt.Println("Before Bubble Sort: ", number_list)
	// fmt.Println("After Bubble Sort: ", sort.BubbleSort(number_list))

	// Cocktail Sort
	// copy(number_list, original_number_list)
	// fmt.Println("Before Cocktail Sort: ", number_list)
	// fmt.Println("After Cocktail Sort: ", sort.CocktailSort(number_list))

	// Comb Sort
	// copy(number_list, original_number_list)
	// fmt.Println("Before Comb Sort: ", number_list)
	// fmt.Println("After Comb Sort: ", sort.CombSort(number_list))

	// Selection Sort
	// copy(number_list, original_number_list)
	// fmt.Println("Before Selection Sort: ", number_list)
	// fmt.Println("After Selection Sort: ", sort.SelectionSort(number_list))

	// Gnome Sort
	// copy(number_list, original_number_list)
	// fmt.Println("Before Gnome Sort: ", number_list)
	// fmt.Println("After Gnome Sort: ", sort.GnomeSort(number_list))

	// Insertion Sort
	// copy(number_list, original_number_list)
	// fmt.Println("Before Insertion Sort: ", number_list)
	// fmt.Println("After Insertion Sort: ", sort.InsertionSort(number_list))

	// Bucket Sort
	// copy(number_list, original_number_list)
	// fmt.Println("Before Bucket Sort: ", number_list)
	// fmt.Println("After Bucket Sort: ", sort.BucketSort(number_list))

	// Shell Sort
	// copy(number_list, original_number_list)
	// fmt.Println("Before Shell Sort: ", number_list)
	// fmt.Println("After Shell Sort: ", sort.ShellSort(number_list))

	// Counting Sort
	// copy(number_list, original_number_list)
	// fmt.Println("Before Counting Sort: ", number_list)
	// fmt.Println("After Counting Sort: ", sort.CountingSort(number_list))

	// Radix Sort
	// copy(number_list, original_number_list)
	// fmt.Println("Before Radix Sort: ", number_list)
	// fmt.Println("After Radix Sort: ", sort.RadixSort(number_list))

	// Singly Linked List
	// l := link_list.LinkedList{}
	// l.Append(1)
	// l.Append(2)
	// l.Append(3)
	// l.Append(2)
	// l.Insert(0)
	// l.Remove(2)
	// l.Print()
}