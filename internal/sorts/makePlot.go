package sorts

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
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

	//line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
	line.SetXAxis(stringArrayX).
		AddSeries("Almost sorted(90/10) case", generateLineAverageItems(arrYAlmost, quantity)).
		AddSeries("Already sorted case", generateLineAverageItems(arrYBest, quantity)).
		AddSeries("Random case", generateLineAverageItems(arrY, quantity)).
		AddSeries("Reverse Sorted case", generateLineAverageItems(arrYWorst, quantity)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: &tr, ConnectNulls: &tr}))
	f, _ := os.Create(fileName)
	_ = line.Render(f)
}
