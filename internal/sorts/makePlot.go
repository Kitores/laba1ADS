package sorts

import (
	"AlgsDataStruct/internal/measurements/regresion"

	//"AlgsDataStruct/internal/measurements"
	"AlgsDataStruct/internal/readCSV"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"strconv"

	"os"
)

// generate data for line chart
func generateLineAverageItems(arrY []int64, quantity int) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < quantity; i++ {
		items = append(items, opts.LineData{Name: "microseconds", Value: arrY[i] / 1e3})
	}
	//fmt.Println(items)
	return items
}

func CreateLineChart(arrX []int, arrY, arrYWorst, arrYBest, arrYAlmost []int64, quantity int, SortName string) {
	fileName := fmt.Sprintf("./internal/sorts/plots/%s.html", SortName)

	// create a new line
	line := charts.NewLine()

	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeWonderland,
			//Theme: types.ThemeRoma,
			//Theme: types.ThemeRomantic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    SortName,
			Subtitle: "This is fun to use!",
		}),
	)
	tr := false
	// Put data into instance

	stringArrayX := make([]string, len(arrX))
	// Преобразуем каждый элемент массива целых чисел в строку
	for i, v := range arrX {
		stringArrayX[i] = strconv.Itoa(v) // Используем strconv.Itoa для преобразования
	}

	line.SetXAxis(stringArrayX).
		AddSeries("Almost sorted(90/10) case", generateLineAverageItems(arrYAlmost, quantity)).
		AddSeries("Already sorted case", generateLineAverageItems(arrYBest, quantity)).
		AddSeries("Random case", generateLineAverageItems(arrY, quantity)).
		AddSeries("Reverse Sorted case", generateLineAverageItems(arrYWorst, quantity)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: &tr, ConnectNulls: &tr}))
	f, _ := os.Create(fileName)
	_ = line.Render(f)
}

func generateLineAverageItemsCSV(arrY []int64, quantity int) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < quantity; i++ {
		items = append(items, opts.LineData{Name: "microseconds", Value: arrY[i] / 1e3})
	}
	//fmt.Println(items)
	return items
}
func CreateLineChartCSV(sortsData []readCSV.Sort, title string, group string) {
	fileName := fmt.Sprintf("./internal/sorts/plots/%s/%s.html", group, title)

	// create a new line
	line := charts.NewLine()

	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeWonderland,
			//Theme: types.ThemeRoma,
			//Theme: types.ThemeRomantic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    title,
			Subtitle: "This is fun to use!",
		}),
	)
	tr := false

	for _, Sort := range sortsData {
		line.SetXAxis(Sort.Sizes).
			AddSeries(fmt.Sprintf("%s", Sort.Name), generateLineAverageItemsCSV(Sort.ExecTimes, len(Sort.Sizes))).
			SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: &tr, ConnectNulls: &tr}))
	}

	f, _ := os.Create(fileName)
	_ = line.Render(f)
}
func CreateLineByPlotterCSV(sortsData []readCSV.Sort, title string, group string) {
	fileName := fmt.Sprintf("./internal/sorts/plots/%s/%s.png", group, title)
	collors := []color.RGBA{{R: 0, G: 0, B: 255, A: 255}, {R: 255, G: 0, B: 0, A: 255}, {R: 0, G: 255, B: 0, A: 255}, {R: 255, G: 127, B: 80, A: 255}, {R: 245, G: 40, B: 150, A: 255}, {R: 106, G: 90, B: 205, A: 255}, {R: 60, G: 179, B: 113, A: 255}}

	plt := plot.New()

	plt.Title.Text = title
	plt.X.Label.Text = "sizes"
	plt.Y.Label.Text = "nanoseconds"
	var sizeInt int
	for i, Sort := range sortsData {
		tmpPoints := make(plotter.XYs, len(Sort.ExecTimes))
		for j, ExecTime := range Sort.ExecTimes {
			tmpPoints[j].Y = float64(ExecTime)
			sizeInt, _ = strconv.Atoi(Sort.Sizes[j])
			tmpPoints[j].X = float64(sizeInt)
		}
		tmpPlot, err := plotter.NewLine(tmpPoints)
		if err != nil {
			panic(err)
		}
		tmpPlot.LineStyle.Color = collors[i]
		plt.Legend.Top = true
		//plt.Legend.Left = true
		plt.Legend.Add(Sort.Name, tmpPlot)
		plt.Add(tmpPlot)
	}
	if err := plt.Save(8*vg.Inch, 4*vg.Inch, fileName); err != nil {
		panic(err)
	}

}

func CreateLineChartByPlotter(arrX []float64, arrY, arrYWorst, arrYBest, arrYAlmost []int64, quantity int, SortName string) {
	fileName := fmt.Sprintf("./internal/sorts/plots/pngPlots/%s.png", SortName)

	plt := plot.New()

	plt.Title.Text = SortName
	plt.X.Label.Text = "sizes"
	plt.Y.Label.Text = "nanoseconds"

	//  random arr
	randomCasePoints := make(plotter.XYs, len(arrX))
	for i := range arrX {
		randomCasePoints[i].X = float64(arrX[i])
		randomCasePoints[i].Y = float64(arrY[i])
	}
	//	reversed arr
	reversedSortedPoints := make(plotter.XYs, len(arrX))
	for i := range arrX {
		reversedSortedPoints[i].X = float64(arrX[i])
		reversedSortedPoints[i].Y = float64(arrYWorst[i])
	}
	//	sorted arr
	alreadySortedPoints := make(plotter.XYs, len(arrX))
	for i := range arrX {
		alreadySortedPoints[i].X = float64(arrX[i])
		alreadySortedPoints[i].Y = float64(arrYBest[i])
	}
	//	almost sorted arr
	almostSortedPoints := make(plotter.XYs, len(arrX))
	for i := range arrX {
		almostSortedPoints[i].X = float64(arrX[i])
		almostSortedPoints[i].Y = float64(arrYAlmost[i])
	}
	averagePlot, err := plotter.NewLine(randomCasePoints)
	if err != nil {
		panic(err)
	}
	worstPlot, err := plotter.NewLine(reversedSortedPoints)
	if err != nil {
		panic(err)
	}
	BestPlot, err := plotter.NewLine(alreadySortedPoints)
	if err != nil {
		panic(err)
	}
	AlmostPlot, err := plotter.NewLine(almostSortedPoints)
	if err != nil {
		panic(err)
	}

	averagePlot.LineStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	worstPlot.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	BestPlot.LineStyle.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	AlmostPlot.LineStyle.Color = color.RGBA{R: 255, G: 127, B: 80, A: 255}
	plt.Add(averagePlot, worstPlot, BestPlot, AlmostPlot)

	plt.Legend.Add("Random Case", averagePlot)
	plt.Legend.Add("Reversed", worstPlot)
	plt.Legend.Add("Sorted", BestPlot)
	plt.Legend.Add("Almost sorted", AlmostPlot)

	plt.X.Tick.Marker = plot.DefaultTicks{}
	plt.Y.Tick.Marker = plot.DefaultTicks{}

	//plt.X.Tick.Marker = plot.ConstantTicks(ticks)

	if err := plt.Save(10*vg.Inch, 5*vg.Inch, fileName); err != nil {
		panic(err)
	}
}

func PolynomialRegressionPlot(x []float64, y []float64, poly *regresion.Polynomial, name, Type string) {
	p := plot.New()

	p.Title.Text = fmt.Sprintf("%s %s Polynomial Regression", name, Type)
	p.X.Label.Text = "sizes"
	p.Y.Label.Text = "nanoseconds"

	//  data points
	realPoints := make(plotter.XYs, len(x))
	for i := range x {
		realPoints[i].X = x[i]
		realPoints[i].Y = y[i]
	}

	realPlot, err := plotter.NewScatter(realPoints)
	if err != nil {
		panic(err)
	}

	//realPlot.GlyphStyle = plotter.ColorBar{Colors: 1}

	// Predicted data points
	newXLen := 1000
	predictedPoints := make(plotter.XYs, newXLen)
	//predictedPoints := make(plotter.XYs, len(x))
	//for i := range x {
	//	predictedPoints[i].X = x[i]
	//	predictedPoints[i].Y = poly.Predict(x[i])
	//}
	newX := linspace(x[0], x[len(x)-1], newXLen)
	for i := range newX {
		predictedPoints[i].X = newX[i]
		predictedPoints[i].Y = poly.Predict(newX[i])
	}

	predictedPlot, err := plotter.NewLine(predictedPoints)
	if err != nil {
		panic(err)
	}

	predictedPlot.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	//ticksX := createTicks(x)
	//ticksY := createTicks(y)
	//p.X.Tick.Marker = plot.ConstantTicks(ticksX)
	//p.Y.Tick.Marker = plot.ConstantTicks(ticksY)

	p.X.Tick.Marker = plot.DefaultTicks{}
	p.Y.Tick.Marker = plot.DefaultTicks{}

	//grid := plotter.NewGrid()
	//grid.Horizontal.Width.Points()
	//grid.Vertical.Width.Points()
	//p.Add(grid)

	p.Add(realPlot, predictedPlot)

	// Save the plot to a PNG file
	if err := p.Save(8*vg.Inch, 4*vg.Inch, fmt.Sprintf("./internal/sorts/plots/regression/%s_polynomial_regression_%s.png", name, Type)); err != nil {
		panic(err)
	}
}

func linspace(start, end float64, num int) []float64 {
	if num <= 0 {
		return []float64{}
	}
	step := (end - start) / float64(num-1)
	result := make([]float64, num)
	for i := 0; i < num; i++ {
		result[i] = start + step*float64(i)
	}
	return result
}

func createTicks(values []float64) []plot.Tick {
	ticks := make([]plot.Tick, len(values))
	for i, value := range values {
		ticks[i] = plot.Tick{Value: value, Label: fmt.Sprintf("%.1f", int(value/1e3))} // Форматируем метку как строку
		//ticks[i] = plot.Tick{Label: fmt.Sprintf("%.1f", value)} // Форматируем метку как строку
	}
	return ticks
}

//func createTicksNorm(lenX []float64) (x, y []plot.Tick) {
//	ticks := make([]plot.Tick, len(values))
//	xx
//	for i, value := range lenX {
//		x[i] = plot.Tick{Value: value, Label: fmt.Sprintf("%.1f", value/1e3)} // Форматируем метку как строку
//		y[i]
//		//ticks[i] = plot.Tick{Label: fmt.Sprintf("%.1f", value)} // Форматируем метку как строку
//	}
//	return x, y
//}
