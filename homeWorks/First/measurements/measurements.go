package measurements

import (
	"AlgsDataStruct/homeWorks/First/sorts"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func TimeMeasurement(sortName string, seed int64) (arrX []int, arrY, arrYWorst, arrYBest []int64) {
	var execTime int64
	var worstExecTime int64
	var bestExecTime int64
	sizes := []int{1000, 5000, 10000, 30000, 50000, 70000, 100000, 200000, 300000}
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
		case "Insertion Sort":
			execTime = measuringInsertionSort(arr)
			bestExecTime = measuringInsertionSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringInsertionSort(arr)
		case "Quick Sort":
			execTime = measuringQuickSort(arr)
			bestExecTime = measuringQuickSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringQuickSort(arr)
		case "Bubble Sort":
			execTime = measuringBubbleSort(arr)
			bestExecTime = measuringBubbleSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringBubbleSort(arr)
		case "Merge Sort":
			execTime = measuringMergeSort(arr)
			bestExecTime = measuringMergeSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringMergeSort(arr)
		case "Shell Sort(shellGaps)":
			execTime = measuringShellSort(arr)
			bestExecTime = measuringShellSort(arr)
			arr = randArray(size, seed)
			sorts.ReverseQuickSort(arr, 0, size-1)
			worstExecTime = measuringShellSort(arr)
		}
		fmt.Println(execTime, worstExecTime, bestExecTime)

		//genCSVfile(sortName, size, execTime)
		arrYBest = append(arrYBest, bestExecTime)
		arrYWorst = append(arrYWorst, worstExecTime)
		arrY = append(arrY, execTime)
		arrX = append(arrX, size)
	}
	genCSVfile(sortName, arrX, arrY)
	//alpha, beta := linearRegression(arrX, arrY)
	//Plot()
	//fmt.Printf("Уравнение регрессии: y = %f + %f * x\n", alpha, beta)

	return arrX, arrY, arrYWorst, arrYBest
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
func genCSVfile(sortName string, sizes []int, execTimes []int64) {
	fileName := fmt.Sprintf("./homeWorks/First/measurements/%sData.csv", sortName)
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
