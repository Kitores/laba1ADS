package measurements

import (
	"AlgsDataStruct/homeWorks/First/sorts"
	"fmt"
	//"gonum.org/v1/gonum/stat"
	"math/rand"
	"time"
)

func TimeMeasurement(sortName string, seed int64) (arrX []int, arrY, arrYWorst []int64) {
	var execTime int64
	var worstExecTime int64
	sizes := []int{1000, 5000, 10000, 30000, 50000, 70000, 100000, 200000, 300000}
	//sizes := []int{70}

	for _, size := range sizes {

		arr := randArray(size, seed)

		switch sortName {
		case "Selection Sort":
			execTime = measuringSelectionSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringSelectionSort(arr)
		case "Insertion Sort":
			execTime = measuringInsertionSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringInsertionSort(arr)
		case "Quick Sort":
			execTime = measuringQuickSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringQuickSort(arr)
		case "Bubble Sort":
			execTime = measuringBubbleSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringBubbleSort(arr)
		case "Merge Sort":
			execTime = measuringMergeSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringMergeSort(arr)
		}
		fmt.Println(execTime, worstExecTime)
		arrYWorst = append(arrYWorst, worstExecTime)
		arrY = append(arrY, execTime)
		arrX = append(arrX, size)
	}
	//alpha, beta := stat.LinearRegression(float64(arrX), arrY, nil, false)
	return arrX, arrY, arrYWorst
}

func TimeMeasurementWorst(sortName string, seed int64) (arrYWorst []int64) {
	var worstExecTime int64
	sizes := []int{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000}
	//sizes := []int{70}

	for _, size := range sizes {

		arr := randArray(size, seed)

		switch sortName {
		case "Selection Sort":
			worstExecTime = measuringSelectionSort(arr)
		case "Insertion Sort":
			worstExecTime = measuringInsertionSort(arr)
		case "Quick Sort":
			worstExecTime = measuringQuickSort(arr)
		case "Bubble Sort":
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringBubbleSort(arr)
		case "Merge Sort":
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringBubbleSort(arr)
		}
		fmt.Println(worstExecTime)
		arrYWorst = append(arrYWorst, worstExecTime)
		//arrX = append(arrX, size)
	}
	return arrYWorst
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

	//os.Stdout()
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
	timeEnd := time.Since(timeStart)
	Time := timeEnd.Nanoseconds()
	return Time
}

func measuringMergeSort(arr []int) int64 {
	timeStart := time.Now()
	sorts.MergeSort(arr)
	timeEnd := time.Now()
	timeFinally := timeEnd.Sub(timeStart)
	Time := timeFinally.Nanoseconds()
	return Time
}
