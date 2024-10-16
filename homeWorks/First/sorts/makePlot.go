package sorts

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"strconv"

	//"math/rand"
	"os"
)

// generate random data for line chart
func generateLineItems(arrY []int, quantity int) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < quantity; i++ {
		y := arrY[i]
		fmt.Println(y)
		items = append(items, opts.LineData{Name: "microseconds", YAxisIndex: y, Value: arrY[i] / 1e3})
	}
	fmt.Println(items)
	return items
}
func CreateLineChart(arrX []int, arrY []int, quantity int, SortName string) {
	fileName := fmt.Sprintf("%s.html", SortName)

	// create a new line instance
	line := charts.NewLine()

	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
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

	//line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
	line.SetXAxis(stringArrayX).
		AddSeries("Category A", generateLineItems(arrY, quantity)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: &tr}))
	f, _ := os.Create(fileName)
	_ = line.Render(f)
}
