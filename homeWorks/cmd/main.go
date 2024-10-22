package main

import (
	"AlgsDataStruct/homeWorks/First/measurements"
	"AlgsDataStruct/homeWorks/First/sorts"
	"fmt"
	"time"
)

//var Sortnames = []string{"Selection Sort", "Insertion Sort", "Quick Sort", "Bubble Sort", "Merge Sort"}

var Sortnames = []string{"Quick Sort"}

func main() {
	//sorts.Try1()
	seed := time.Now().UnixNano()
	//seed := int64(12)
	for _, name := range Sortnames {
		arrX, arrY, arrYWorst, arrYBest := measurements.TimeMeasurement(name, seed)
		//arrYWorst := measurements.TimeMeasurementWorst(name, seed)
		quantity := len(arrX)
		fmt.Println(name)
		sorts.CreateLineChart(arrX, arrY, arrYWorst, arrYBest, quantity, name)
	}
}
