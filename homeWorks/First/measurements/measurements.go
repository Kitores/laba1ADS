package measurements

import (
	"AlgsDataStruct/homeWorks/First/sorts"
	"math/rand"
	"time"
)

func TimeMeasurement(sortName string, seed int64) (arrX []int, arrY []int64) {
	var execTime int64
	sizes := []int{1000, 5000, 10000, 30000, 50000, 70000, 100000}

	for _, size := range sizes {

		arr := randArray(size, seed)

		switch sortName {
		case "Selection Sort":
			execTime = measuringSelectionSort(arr)
		case "Insertion Sort":
			execTime = measuringInsertionSort(arr)
		case "Quick Sort":
			execTime = measuringQuickSort(arr)
		case "Bubble Sort":
			execTime = measuringBubbleSort(arr)
		}

		arrY = append(arrY, execTime)
		arrX = append(arrX, size)
	}

	return arrX, arrY
}
func randArray(size int, seed int64) []int {
	rand.Seed(seed)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100000)
	}
	return arr
}
func measuringSelectionSort(arr []int) int64 {
	timeStart := time.Now()
	sorts.SelectSort(arr)
	timeEnd := time.Now()
	timeFinally := timeEnd.Sub(timeStart)
	Time := timeFinally.Nanoseconds()
	return Time
}

func measuringInsertionSort(arr []int) int64 {
	timeStart := time.Now()
	sorts.InsertionSort(arr)
	timeEnd := time.Now()
	timeFinally := timeEnd.Sub(timeStart)
	Time := timeFinally.Nanoseconds()
	return Time
}

func measuringQuickSort(arr []int) int64 {
	timeStart := time.Now()
	sorts.QuickSort(arr, 0, len(arr)-1)
	timeEnd := time.Now()
	timeFinally := timeEnd.Sub(timeStart)
	Time := timeFinally.Nanoseconds()
	return Time
}

func measuringBubbleSort(arr []int) int64 {
	timeStart := time.Now()
	sorts.BubbleSort(arr)
	timeEnd := time.Now()
	timeFinally := timeEnd.Sub(timeStart)
	Time := timeFinally.Nanoseconds()
	return Time
}
