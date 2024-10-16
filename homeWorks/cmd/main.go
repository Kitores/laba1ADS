package main

import (
	"AlgsDataStruct/homeWorks/First/measurements"
	"AlgsDataStruct/homeWorks/First/sorts"
	"fmt"
)

var Sortnames = []string{"Selection Sort", "Insertion Sort", "Quick Sort"}

func main() {
	//sorts.Try1()
	//seed := time.Now().UnixNano()
	seed := int64(12)
	for _, name := range Sortnames {
		arrX, arrY := measurements.TimeMeasurement(name, seed)
		quantity := len(arrX)
		fmt.Println(name)
		sorts.CreateLineChart(arrX, arrY, quantity, name)
	}

}
