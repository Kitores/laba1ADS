package main

import (
	"AlgsDataStruct/internal/measurements"
	"AlgsDataStruct/internal/sorts"
	"fmt"
	"time"
)

var SortNames = []string{"Selection Sort", "Insertion Sort", "Quick Sort", "Bubble Sort", "Merge Sort", "Shell Sort(shellGaps)", "Shell Sort(hibbardGaps)", "Shell Sort(prattGaps)", "Heap Sort"}

//var SortNames = []string{"Heap Sort"}

func main() {
	//sorts.Test()
	seed := time.Now().UnixNano()
	//seed := int64(12)
	for _, name := range SortNames {
		arrX, arrY, arrYWorst, arrYBest, arrYAlmost := measurements.TimeMeasurement(name, seed)
		//arrYWorst := measurements.TimeMeasurementWorst(name, seed)
		quantity := len(arrX)
		fmt.Println(name)
		sorts.CreateLineChart(arrX, arrY, arrYWorst, arrYBest, arrYAlmost, quantity, name)
	}
}
