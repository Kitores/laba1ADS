package measurements

import (
	"AlgsDataStruct/internal/sorts"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func TimeMeasurement(sortName string, seed int64) (arrX []int, arrY, arrYWorst, arrYBest, arrYAlmost []int64) {
	var execTime int64
	var worstExecTime int64
	var bestExecTime int64
	var almostExecTime int64
	sortsQuadratic := []string{"Selection Sort", "Insertion Sort", "Bubble Sort", "Shell Sort(shellGaps)", "Shell Sort(hibbardGaps)", "Shell Sort(prattGaps)"}
	sizes := []int{1000, 5000, 10000, 30000, 50000, 70000, 100000, 200000, 300000, 500000}
	// O(n^2)
	if contains(sortsQuadratic, sortName) {
		sizes = sizes[:6]
	}
	//sizes := []int{70}

	for _, size := range sizes {

		arr := randArray(size, seed)

		switch sortName {
		case "Selection Sort":
			execTime = measuringSelectionSort(arr)
			bestExecTime = measuringSelectionSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringSelectionSort(arr)
			arr = randArray(size, seed)
			sorts.QuickSort(arr, 0, size*9/10-1)
			almostExecTime = measuringSelectionSort(arr)
		case "Insertion Sort":
			execTime = measuringInsertionSort(arr)
			bestExecTime = measuringInsertionSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringInsertionSort(arr)
			arr = randArray(size, seed)
			sorts.QuickSort(arr, 0, size*9/10-1)
			almostExecTime = measuringInsertionSort(arr)

		case "Quick Sort":
			execTime = measuringQuickSort(arr)
			bestExecTime = measuringQuickSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringQuickSort(arr)
			arr = randArray(size, seed)
			sorts.QuickSort(arr, 0, size*9/10-1)
			almostExecTime = measuringQuickSort(arr)
		case "Bubble Sort":
			execTime = measuringBubbleSort(arr)
			bestExecTime = measuringBubbleSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringBubbleSort(arr)
			arr = randArray(size, seed)
			sorts.QuickSort(arr, 0, size*9/10-1)
			almostExecTime = measuringBubbleSort(arr)
		case "Merge Sort":
			execTime = measuringMergeSort(arr)
			bestExecTime = measuringMergeSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringMergeSort(arr)
			arr = randArray(size, seed)
			sorts.QuickSort(arr, 0, size*9/10-1)
			almostExecTime = measuringMergeSort(arr)
		case "Shell Sort(shellGaps)":
			execTime = measuringShellSort(arr)
			bestExecTime = measuringShellSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringShellSort(arr)
			arr = randArray(size, seed)
			sorts.QuickSort(arr, 0, size*9/10-1)
			almostExecTime = measuringShellSort(arr)
		case "Shell Sort(hibbardGaps)":
			execTime = measuringShellSortHibbard(arr)
			bestExecTime = measuringShellSortHibbard(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringShellSortHibbard(arr)
			arr = randArray(size, seed)
			sorts.QuickSort(arr, 0, size*9/10-1)
			almostExecTime = measuringShellSortHibbard(arr)
		case "Shell Sort(prattGaps)":
			execTime = measuringShellSortPratt(arr)
			bestExecTime = measuringShellSortPratt(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringShellSortPratt(arr)
			arr = randArray(size, seed)
			sorts.QuickSort(arr, 0, size*9/10-1)
			almostExecTime = measuringShellSortPratt(arr)
		case "Heap Sort":
			execTime = measuringHeapSort(arr)
			bestExecTime = measuringHeapSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringHeapSort(arr)
			arr = randArray(size, seed)
			sorts.QuickSort(arr, 0, size*9/10-1)
			almostExecTime = measuringHeapSort(arr)
		}
		fmt.Println(execTime, worstExecTime, bestExecTime, almostExecTime)

		//genCSVfile(sortName, size, execTime)
		arrYAlmost = append(arrYAlmost, almostExecTime)
		arrYBest = append(arrYBest, bestExecTime)
		arrYWorst = append(arrYWorst, worstExecTime)
		arrY = append(arrY, execTime)
		arrX = append(arrX, size)
	}
	//genCSVfile(sortName+"Average", arrX, arrY)
	//genCSVfile(sortName+"Best", arrX, arrYBest)
	//genCSVfile(sortName+"Worst", arrX, arrYWorst)
	//genCSVfile(sortName+"Almost", arrX, arrYAlmost)

	return arrX, arrY, arrYWorst, arrYBest, arrYAlmost
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
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
func measuringShellSort(arr []int) int64 {
	timeStart := time.Now()
	sorts.ShellSort(arr)
	timeEnd := time.Now()
	timeFinally := timeEnd.Sub(timeStart)
	Time := timeFinally.Nanoseconds()
	return Time
}
func measuringShellSortHibbard(arr []int) int64 {
	timeStart := time.Now()
	sorts.ShellSortHibbard(arr)
	timeEnd := time.Now()
	timeFinally := timeEnd.Sub(timeStart)
	Time := timeFinally.Nanoseconds()
	return Time
}
func measuringShellSortPratt(arr []int) int64 {
	timeStart := time.Now()
	sorts.ShellSortPratt(arr)
	timeEnd := time.Now()
	timeFinally := timeEnd.Sub(timeStart)
	Time := timeFinally.Nanoseconds()
	return Time
}
func measuringHeapSort(arr []int) int64 {
	timeStart := time.Now()
	sorts.HeapSort(arr)
	timeEnd := time.Now()
	timeFinally := timeEnd.Sub(timeStart)
	Time := timeFinally.Nanoseconds()
	return Time
}
func genCSVfile(sortName string, sizes []int, execTimes []int64) {
	fileName := fmt.Sprintf("./internal/sorts/csvData/%sData.csv", sortName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	rec := []string{"sizes", "execTimes"}
	err = csvWriter.Write(rec)
	if err != nil {
		fmt.Println(err)
	}
	// Очищаем буфер перед закрытием
	defer csvWriter.Flush()

	for i := 0; i < len(sizes); i++ {
		data := []string{strconv.Itoa(sizes[i]), strconv.Itoa(int(execTimes[i]))}
		if err = csvWriter.Write(data); err != nil {
			fmt.Println(err)
		}
	}

}
