package main

import (
	"AlgsDataStruct/internal/measurements"
	"AlgsDataStruct/internal/readCSV"
	"AlgsDataStruct/internal/sorts"
	"fmt"
	"time"
)

var SortNames = []string{"Selection Sort", "Insertion Sort", "Quick Sort", "Bubble Sort", "Merge Sort", "Shell Sort(shellGaps)", "Shell Sort(hibbardGaps)", "Shell Sort(prattGaps)", "Heap Sort"}

// test
// var SortNames = []string{"Selection Sort"}
var N2Sorts = []string{"Selection Sort", "Insertion Sort", "Bubble Sort"}

var nLogNSorts = []string{"Quick Sort", "Merge Sort", "Shell Sort(shellGaps)", "Shell Sort(hibbardGaps)", "Shell Sort(prattGaps)", "Heap Sort"}

func main() {
	//sorts.Test()
	seed := time.Now().UnixNano()
	//seed := int64(12)
	for _, name := range SortNames {
		arrX, arrY, arrYWorst, arrYBest, arrYAlmost := measurements.TimeMeasurement(name, seed)
		quantity := len(arrX)
		fmt.Println(arrX, arrY, arrYWorst, arrYBest, arrYAlmost, quantity, name)
		sorts.CreateLineChart(arrX, arrY, arrYWorst, arrYBest, arrYAlmost, quantity, name)
	}

	titleNlogN := "nLogN"
	titleN2 := "n^2"
	var SortsDataNLogn = []readCSV.Sort{}
	var SortsDataN2 = []readCSV.Sort{}

	for _, name := range nLogNSorts {
		sortDataFromCSV := readCSV.ReadCSV(name)
		SortsDataNLogn = append(SortsDataNLogn, sortDataFromCSV)
	}
	sorts.CreateLineChartCSV(SortsDataNLogn, titleNlogN)

	for _, name := range N2Sorts {
		sortDataFromCSV := readCSV.ReadCSV(name)
		SortsDataN2 = append(SortsDataN2, sortDataFromCSV)
	}
	sorts.CreateLineChartCSV(SortsDataN2, titleN2)
}
