package sorts

import (
	"fmt"
	"math/rand"
	"time"
)

func Test() {

	//size := []int{100, 500, 1000, 5000, 10000}
	arr := randArray(1000, 2)
	fmt.Println(arr)
	timeStart := time.Now()
	//QuickSort(arr, 0, len(arr)*9/10-1)
	HeapSort(arr)
	timeEnd := time.Since(timeStart)

	fmt.Println(arr, timeEnd)

	//nums := []int{100, 200, 300, 400, 500, 600, 700, 800, 900}
	//nums = nums[:5]
	//fmt.Println(nums)

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
		arr[i] = rand.Intn(1000)
	}
	return arr
}

func StrangeSort(nums []int) []int {
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
	flag := true
	for i := 0; i < N-1 && flag; i++ {
		flag = false
		for j := 0; j < N-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j, j+1)
				flag = true
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

func shellGaps1(length int) []int {
	var gaps = []int{}
	i := length
	for i > 1 {
		i /= 2
		gaps = append(gaps, i)
	}
	return gaps
}

func shellGaps2(n int) []int {
	gaps := []int{}
	for gap := n / 2; gap > 0; gap /= 2 {
		gaps = append(gaps, gap)
	}
	return gaps
}

func hibbardGaps(n int) []int {
	gaps := []int{}
	for i := 1; (1<<i)-1 < n; i++ {
		gaps = append(gaps, (1<<i)-1)
	}
	return gaps
}

func prattGaps(n int) []int {
	gaps := []int{}
	for i := 0; (1 << i) <= n; i++ {
		for j := 0; (1 << j) <= n; j++ {
			gap := (1 << i) * (3 << j)
			if gap <= n {
				gaps = append(gaps, gap)
			}
		}
	}
	return gaps
}

func ShellSort(arr []int) {
	gaps := shellGaps2(len(arr))
	k := len(gaps)
	gap := gaps[0]
	for k >= 1 {
		for i := gap; i < len(arr); i++ {
			a := arr[i]
			for j := i - gap; j >= 0 && arr[j] > a; j -= gap {
				arr[j+gap] = arr[j]
				arr[j] = a
			}
		}
		k--
		gap = gaps[k]
	}
}
func ShellSortHibbard(arr []int) {
	gaps := hibbardGaps(len(arr))
	k := len(gaps)
	gap := gaps[0]
	for k >= 1 {
		for i := gap; i < len(arr); i++ {
			a := arr[i]
			for j := i - gap; j >= 0 && arr[j] > a; j -= gap {
				arr[j+gap] = arr[j]
				arr[j] = a
			}
		}
		k--
		gap = gaps[k]
	}
}
func ShellSortPratt(arr []int) {
	gaps := prattGaps(len(arr))
	k := len(gaps)
	gap := gaps[0]
	for k >= 1 {
		for i := gap; i < len(arr); i++ {
			a := arr[i]
			for j := i - gap; j >= 0 && arr[j] > a; j -= gap {
				arr[j+gap] = arr[j]
				arr[j] = a
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

func heapify(arr []int, i, heapLen int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < heapLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < heapLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		Swap(arr, i, largest)
		heapify(arr, largest, heapLen)
	}

}

func HeapSort(arr []int) {
	heapLen := len(arr)
	//1. make haep
	for i := heapLen/2 - 1; i >= 0; i-- {
		heapify(arr, i, heapLen)
	}
	//2. delete elements
	for i := heapLen - 1; i > 0; i-- {
		Swap(arr, 0, i)
		heapify(arr, 0, i)
	}
}
