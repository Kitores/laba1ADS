package regresion

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

type Polynomial struct {
	Coefficients []float64
}

func LinearRegression(arrX []int, arrY []int64, degree int) mat.Matrix {
	n := len(arrX)
	var arrYfloat64 []float64
	vandermonde := mat.NewDense(n, degree+1, nil)
	for i := 0; i < n; i++ {
		arrYfloat64 = append(arrYfloat64, float64(arrY[i]/1e9))
		for j := 0; j <= degree; j++ {
			vandermonde.Set(i, j, math.Pow(float64(arrX[i]), float64(j)))
		}
	}
	yVector := mat.NewVecDense(n, arrYfloat64)
	fmt.Println(yVector)
	coefficients := mat.NewVecDense(n, nil)
	rowsCOeff, collsCOeff := coefficients.Dims()
	rowsYVec, collsYVec := yVector.Dims()
	fmt.Println(rowsCOeff, collsCOeff, rowsYVec, collsYVec)

	if err := vandermonde.Solve(coefficients, yVector); err != nil {
		fmt.Println("Ошибка при решении уравнения ", err)
		return nil
	}
	fmt.Println("Коэффициенты полиномиальной регрессии:")
	for i := 0; i < coefficients.Len(); i++ {
		fmt.Printf("a%d: %v\n", i, coefficients.At(i, 0))
	}
	return coefficients
}

func (p *Polynomial) Fit(x []float64, y []int64, degree int) {
	n := len(x)
	var yFloat64 []float64
	//Make vandermonde matrix
	Vandermonde := mat.NewDense(n, degree+1, nil)
	for i := 0; i < n; i++ {
		yFloat64 = append(yFloat64, float64(y[i]))
		for j := 0; j <= degree; j++ {
			Vandermonde.Set(i, j, math.Pow(x[i], float64(j)))
		}
	}

	yVector := mat.NewVecDense(n, yFloat64)

	var coeffs mat.VecDense
	var ATA mat.Dense
	var ATY mat.VecDense

	ATA.Mul(Vandermonde.T(), Vandermonde)
	ATY.MulVec(Vandermonde.T(), yVector)

	_ = coeffs.SolveVec(&ATA, &ATY)

	p.Coefficients = coeffs.RawVector().Data

	fmt.Printf("Polynomial = %.2f * x^2 + %.2f * x + %.2f\n", p.Coefficients[2], p.Coefficients[1], p.Coefficients[0])
}

func (p *Polynomial) Predict(x float64) float64 {
	result := 0.0
	for i, coeff := range p.Coefficients {
		result += coeff * math.Pow(x, float64(i))
	}
	return result
}

// // PlotGraph plots the real and predicted values.

//
//// PlotGraph plots the real and predicted values using go-echarts.
//func PlotGraph(x []float64, y []float64, poly *Polynomial) {
//	line := charts.NewLine()
//
//	// Set global options for the chart
//	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Polynomial Regression", Subtitle: " vs Predicted"}))
//
//	// Prepare data for real values
//	realData := make([]opts.LineData, len(x))
//	for i := range x {
//		realData[i] = opts.LineData{Value: y[i]}
//	}
//
//	// Prepare data for predicted values
//	predictedData := make([]opts.LineData, len(x))
//	for i := range x {
//		predictedData[i] = opts.LineData{Value: poly.Predict(x[i])}
//	}
//
//	// Set X axis and add series
//	line.SetXAxis(x).AddSeries(" Values", realData).AddSeries("Predicted Values", predictedData)
//
//	// Save to HTML file
//	f, err := os.Create("polynomial_regression.html")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//
//	if err := line.Render(f); err != nil {
//		panic(err)
//	}
//}
//
//func PlotGraph2(x []float64, y []float64, poly *Polynomial) {
//	scatter := charts.NewScatter()
//	line := charts.NewLine()
//	// Set global options for the chart
//	scatter.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Polynomial Regression", Subtitle: " vs Predicted"}))
//
//	// Prepare data for real values
//	realData := make([]opts.ScatterData, len(x))
//	for i := range x {
//		realData[i] = opts.ScatterData{Value: []interface{}{x[i], y[i]}}
//	}
//
//	// Prepare data for predicted values
//	predictedData := make([]opts.LineData, len(x))
//	for i := range x {
//		predictedData[i] = opts.LineData{Value: poly.Predict(x[i])}
//	}
//
//	// Set X axis and add series
//	scatter.AddSeries(" Values", realData)
//	line.SetXAxis(x).AddSeries("Predicted Values", predictedData)
//	// Save to HTML file
//	f, err := os.Create("polynomial_regression_scatter.html")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//
//	line.Render(f)
//	scatter.Render(f)
//}
//
//func PlotGraph3(x []float64, y []float64, poly *Polynomial) {
//	line := charts.NewLine()
//
//	// Устанавливаем глобальные параметры для графика
//	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Polynomial Regression", Subtitle: "Real vs Predicted"}))
//
//	// Подготавливаем данные для предсказанных значений
//	predictedData := make([]opts.LineData, len(x))
//	for i := range x {
//		predictedData[i] = opts.LineData{Value: poly.Predict(x[i])}
//	}
//
//	// Добавляем серию предсказанных значений в виде линии
//	line.SetXAxis(x).AddSeries("Predicted Values", predictedData)
//
//	// Создаем график разброса для реальных данных
//	scatter := charts.NewScatter()
//	realData := make([]opts.ScatterData, len(x))
//	for i := range x {
//		realData[i] = opts.ScatterData{Value: []interface{}{x[i], y[i]}}
//	}
//
//	// Добавляем серию реальных значений в виде точек
//	scatter.AddSeries("Real Values", realData)
//
//	// Сохраняем график в HTML файл
//	f, err := os.Create("polynomial_regression_combined.html")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//
//	// Рендерим оба графика в один файл
//	line.Render(f)
//	scatter.Render(f)
//}
//
//func PlotGraph4(x []float64, y []float64, poly *Polynomial) {
//	// Создаем новый график
//	line := charts.NewLine()
//
//	// Устанавливаем глобальные параметры для графика
//	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Polynomial Regression", Subtitle: "Real vs Predicted"}))
//
//	// Подготавливаем данные для предсказанных значений
//	predictedData := make([]opts.LineData, len(x))
//	for i := range x {
//		predictedData[i] = opts.LineData{Value: poly.Predict(x[i])}
//	}
//
//	// Добавляем серию предсказанных значений в виде линии
//	line.SetXAxis(x).AddSeries("Predicted Values", predictedData)
//
//	// Создаем график разброса для реальных данных
//	realData := make([]opts.ScatterData, len(x))
//	for i := range x {
//		realData[i] = opts.ScatterData{Value: []interface{}{x[i], y[i]}}
//	}
//
//	// Добавляем серию реальных значений в виде точек
//	line.AddSeries("Real Values", realData)
//
//	// Сохраняем график в HTML файл
//	f, err := os.Create("polynomial_regression_combined.html")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//
//	if err := line.Render(f); err != nil {
//		panic(err)
//	}
//}
//func PlotGraph5(x []float64, y []float64, poly *Polynomial) {
//	// Создаем новый график
//	chart := charts.NewLine()
//
//	// Устанавливаем глобальные параметры для графика
//	chart.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Polynomial Regression", Subtitle: "Real vs Predicted"}))
//
//	// Подготавливаем данные для предсказанных значений
//	predictedData := make([]opts.LineData, len(x))
//	for i := range x {
//		predictedData[i] = opts.LineData{Value: poly.Predict(x[i])}
//	}
//
//	// Добавляем серию предсказанных значений в виде линии
//	chart.SetXAxis(x).AddSeries("Predicted Values", predictedData)
//
//	// Создаем график разброса для реальных данных
//	realData := make([]opts.ScatterData, len(x))
//	for i := range x {
//		realData[i] = opts.ScatterData{Value: []interface{}{x[i], y[i]}}
//	}
//
//	// Добавляем серию реальных значений в виде точек
//	chart.AddSeries("Real Values", realData)
//
//	// Сохраняем график в HTML файл
//	f, err := os.Create("polynomial_regression_combined.html")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//
//	if err := chart.Render(f); err != nil {
//		panic(err)
//	}
//}
