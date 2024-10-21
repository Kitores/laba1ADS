package sorts

import (
	"fmt"
	"math/rand"
	"time"
)

func Try1() {
	//arr := []int{-9, -231, 67, -50, 42, -6, 4, 54, 2, 7, 71, 34, 242, 533, 23352, 523, 13, 11}
	//size := []int{100, 500, 1000, 5000, 10000}
	arr := randArray(10, 2)
	fmt.Println(arr)
	timeStart := time.Now()
	srt := MergeSort(arr)
	timeEnd := time.Since(timeStart)

	fmt.Println(srt, timeEnd)
	//
	//worstArr := randArray(size[4], 2)

	//CrazySort(arr)
	//SelectSort(arr)
	//InsertionSort(arr)

	//InsertionSort(arr)
	//ShellSort(arr, Gaps(lenght))
	//QuickSort(arr, 0, len(arr)-1)
	//BubbleSort(arr)

	s := time.Now().UnixNano()
	testarr := randArray(10, s)
	testarr1 := randArray(10, s)
	testarr2 := randArray(10, s)
	fmt.Println(testarr, testarr1, testarr2)

}

func randArray(size int, seed int64) []int {
	rand.Seed(seed)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(10000)
	}
	return arr
}

func CrazySort(nums []int) []int {
	counter := 0
	for {
		if IsSorted(nums) {
			fmt.Println(counter)
			return nums
		} else {
			for i := 0; i < len(nums)-1; i++ {
				ok := nums[i] < nums[i+1]
				if !ok {
					for ; i < len(nums)-1; i++ {
						Swap(nums, rand.Intn(len(nums)), i)
					}
				}
			}
		}
		counter++
	}
	return nums
}

func BubbleSort(arr []int) {
	N := len(arr)
	for i := 0; i < N-1; i++ {
		for j := 0; j < N-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j, j+1)
			}
		}
	}
}

func Swap(arr []int, a, b int) []int {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
	return arr
}
func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}
func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j := i - 1
		a := nums[i]
		for j >= 0 && nums[j] > a {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = a
	}
}

func SelectSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		minindex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minindex] {
				minindex = j
			}
		}
		if minindex != i {
			nums[i], nums[minindex] = nums[minindex], nums[i]
		}
	}
}

func ShelliSort(nums []int) {

	for i := 1; i < len(nums); i++ {
		a := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > a {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = a
	}
}
func Gaps(length int) []int {
	var gaps = []int{}
	i := length
	for i > 1 {
		i /= 2
		gaps = append(gaps, i)
	}
	return gaps
}

func ShellSort(nums []int, gaps []int) {
	k := len(gaps)
	gap := gaps[0]
	for k >= 1 {
		for i := gap; i < len(nums); i++ {
			a := nums[i]
			for j := i - gap; j >= 0 && nums[j] > a; j -= gap {
				nums[j+gap] = nums[j]
				nums[j] = a
			}
		}
		k--
		gap = gaps[k]
	}
}
func partition(arr []int, left, right int) int {
	pivot := arr[right]
	j := left - 1
	for i := left; i <= right; i++ {
		if arr[i] < pivot {
			j++
			Swap(arr, i, j)
		}
	}
	Swap(arr, j+1, right)
	return j + 1
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		QuickSort(arr, left, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, right)
	}
}

//	func MergeSort(arr []int) []int {
//		var left = arr[0]
//		var right = arr[len(arr)-1]
//		var mid = arr[len(arr)/2]
//
//		for i := 1; i < len(arr)-1; i++ {
//
//		}
//	}
func reversePartition(arr []int, left, right int) int {
	pivot := arr[right]
	j := left - 1
	for i := left; i <= right; i++ {
		if arr[i] > pivot {
			j++
			Swap(arr, i, j)
		}
	}
	Swap(arr, j+1, right)
	return j + 1
}

func ReverseQuickSort(arr []int, left, right int) {
	if left < right {
		pivotIndex := reversePartition(arr, left, right)
		ReverseQuickSort(arr, left, pivotIndex-1)
		ReverseQuickSort(arr, pivotIndex+1, right)
	}
}

func merge(left, right []int) []int {
	arr := []int{}
	var r, l = 0, 0
	for len(left)-l > 0 && len(right)-r > 0 {
		if left[l] < right[r] {
			arr = append(arr, left[l])
			l++
		} else {
			arr = append(arr, right[r])
			r++
		}
	}
	for l < len(left) {
		arr = append(arr, left[l])
		l++
	}
	for r < len(right) {
		arr = append(arr, right[r])
		r++
	}
	//fmt.Println(arr, "Merge")
	return arr
}

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	var L, R []int
	half := length / 2
	if length%2 == 0 {
		L = make([]int, half)
		R = make([]int, half)
	} else {
		L = make([]int, half)
		R = make([]int, half+1)
	}

	for i := 0; i < half; i++ {
		L[i] = arr[i]
	}
	for j := 0; j < half+length%2; j++ {
		R[j] = arr[j+half]
	}
	//fmt.Println(L, R)
	return merge(MergeSort(L), MergeSort(R))

}

//func merge(left, right []int) []int {
//	arr := []int{}
//	i, j := 0, 0
//
//	for i < len(left) && j < len(right) {
//		if left[i] < right[j] {
//			arr = append(arr, left[i])
//			i++
//		} else {
//			arr = append(arr, right[j])
//			j++
//		}
//	}
//
//	// Копируем оставшиеся элементы (если есть)
//	for i < len(left) {
//		arr = append(arr, left[i])
//		i++
//	}
//
//	for j < len(right) {
//		arr = append(arr, right[j])
//		j++
//	}
//
//	return arr
//}
//
//// Функция сортировки слиянием
//func MergeSort(arr []int) []int {
//	if len(arr) <= 1 {
//		return arr
//	}
//
//	mid := len(arr) / 2
//	left := MergeSort(arr[:mid])
//	right := MergeSort(arr[mid:])
//
//	return merge(left, right)
//}
