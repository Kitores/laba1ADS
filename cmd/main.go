package main

import (
	"AlgsDataStruct/internal/makePlots"
	"AlgsDataStruct/internal/measurements"
	"AlgsDataStruct/internal/readCSV"
	"AlgsDataStruct/internal/sorts"
	"fmt"
	"time"
)

var SortNames = []string{"Bubble Sort", "Selection Sort", "Insertion Sort", "Shell Sort(shellGaps)", "Shell Sort(hibbardGaps)", "Shell Sort(prattGaps)", "Quick Sort", "Merge Sort", "Heap Sort"}

// test
// var SortNames = []string{"Merge Sort"}
var N2Sorts = []string{"Selection Sort", "Insertion Sort", "Bubble Sort"}

var nLogNSorts = []string{"Quick Sort", "Merge Sort", "Shell Sort(shellGaps)", "Shell Sort(hibbardGaps)", "Shell Sort(prattGaps)", "Heap Sort"}

//var types = []string{"Average", "Worst", "Best", "Almost"}

func main() {
	sorts.Test()

	seed := time.Now().UnixNano()
	//seed := int64(12)
	var arrX []float64
	var arrY, arrYWorst, arrYBest, arrYAlmost []int64
	for _, name := range SortNames {
		fmt.Println(name)
		arrX, arrY, arrYWorst, arrYBest, arrYAlmost = measurements.TimeMeasurement(name, seed)
		quantity := len(arrX)

		makePlots.GenRegressionPlot(name, arrX, arrY, arrYWorst, arrYBest, arrYAlmost)
		//fmt.Println(arrX, arrY, arrYWorst, arrYBest, arrYAlmost, quantity, name)
		sorts.CreateLineChartByPlotter(arrX, arrY, arrYWorst, arrYBest, arrYAlmost, quantity, name)
		////sorts.CreateLineChart(pkg.Float64ToInt(arrX), arrY, arrYWorst, arrYBest, arrYAlmost, quantity, name)
	}

	titleNlogN := "nLogN"
	titleN2 := "n^2"
	groupN2 := "n2Plot"
	groupNLogn := "nlognPlot"
	var sortDataFromCSV readCSV.Sort
	var SortsDataNLogn = []readCSV.Sort{}
	var SortsDataN2 = []readCSV.Sort{}

	for _, name := range nLogNSorts {
		sortDataFromCSV = readCSV.ReadCSV(name)
		SortsDataNLogn = append(SortsDataNLogn, sortDataFromCSV)
	}
	//sorts.CreateLineChartCSV(SortsDataNLogn, titleNlogN, groupNLogn)
	sorts.CreateLineByPlotterCSV(SortsDataNLogn, titleNlogN, groupNLogn)

	for _, name := range N2Sorts {
		sortDataFromCSV = readCSV.ReadCSV(name)
		SortsDataN2 = append(SortsDataN2, sortDataFromCSV)
	}
	//sorts.CreateLineChartCSV(SortsDataN2, titleN2, groupN2)
	sorts.CreateLineByPlotterCSV(SortsDataN2, titleN2, groupN2)
}
