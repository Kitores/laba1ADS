package makePlots

import (
	"AlgsDataStruct/internal/measurements/regresion"
	"AlgsDataStruct/internal/sorts"
	"AlgsDataStruct/pkg"
	"fmt"
)

var types = []string{"Average", "Worst", "Best", "Almost"}

func GenRegressionPlot(name string, arrX []float64, arrY, arrYWorst, arrYBest, arrYAlmost []int64) {
	//arrX := pkg.IntToFloat64(arrX)

	//average
	fmt.Print("Average ")
	averagePoly := &regresion.Polynomial{}
	averagePoly.Fit(arrX, arrY, 2)
	//worst
	fmt.Print("Worst ")
	worstPoly := &regresion.Polynomial{}
	worstPoly.Fit(arrX, arrYWorst, 2)
	//best
	fmt.Print("Best ")
	bestPoly := &regresion.Polynomial{}
	bestPoly.Fit(arrX, arrYBest, 2)
	//almost
	fmt.Print("Almost ")
	almostPoly := &regresion.Polynomial{}
	almostPoly.Fit(arrX, arrYAlmost, 2)

	// Predict values and print results
	for _, val := range arrX {
		//average
		averagePoly.Predict(val)
		//worst
		worstPoly.Predict(val)
		//best
		bestPoly.Predict(val)
		//almost
		almostPoly.Predict(val)
	}

	//average
	sorts.PolynomialRegressionPlot(arrX, pkg.Int64ToFloat64(arrY), averagePoly, name, types[0])
	//worst
	sorts.PolynomialRegressionPlot(arrX, pkg.Int64ToFloat64(arrYWorst), worstPoly, name, types[1])
	//best
	sorts.PolynomialRegressionPlot(arrX, pkg.Int64ToFloat64(arrYBest), bestPoly, name, types[2])
	//almost
	sorts.PolynomialRegressionPlot(arrX, pkg.Int64ToFloat64(arrYAlmost), almostPoly, name, types[3])
}
